package commands

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewAuthCmd creates the "auth" command group.
func NewAuthCmd(envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with the Aula platform (login, logout, status)",
	}

	// login
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Log in via browser-based OIDC flow (UniLogin or MitID)",
		Run: func(cmd *cobra.Command, args []string) {
			level, _ := cmd.Flags().GetInt("level")
			handleLogin(level, *envFlag)
		},
	}
	loginCmd.Flags().Int("level", 2, "Authentication level: 2 for UniLogin, 3 for MitID")
	cmd.AddCommand(loginCmd)

	// logout
	cmd.AddCommand(&cobra.Command{
		Use:   "logout",
		Short: "Log out and clear the current session",
		Run: func(cmd *cobra.Command, args []string) {
			handleLogout(*envFlag)
		},
	})

	// status
	cmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Show current authentication status",
		Run: func(cmd *cobra.Command, args []string) {
			handleStatus()
		},
	})

	// refresh
	cmd.AddCommand(&cobra.Command{
		Use:   "refresh",
		Short: "Force a token refresh using the stored refresh token",
		Run: func(cmd *cobra.Command, args []string) {
			handleRefresh(*envFlag)
		},
	})

	return cmd
}

func parseAuthLevel(level int) (aulaapi.AuthLevel, error) {
	switch level {
	case 2:
		return aulaapi.AuthLevel2, nil
	case 3:
		return aulaapi.AuthLevel3, nil
	default:
		return "", fmt.Errorf("invalid auth level %d: must be 2 or 3", level)
	}
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}

func handleLogin(level int, envOverride string) {
	authLevel, err := parseAuthLevel(level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	environment := cli.ResolveEnvironment(envOverride)
	endpoints := aulaapi.OidcEndpointsForEnvironment(environment)
	pkce, err := aulaapi.GeneratePkce()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	state, err := aulaapi.GenerateState()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	params := &aulaapi.AuthorizeParams{
		AuthLevel:     authLevel,
		CodeChallenge: pkce.CodeChallenge,
		State:         state,
		RedirectURI:   "", // uses default
	}
	authorizeURL := aulaapi.BuildAuthorizeURL(&endpoints, params)

	fmt.Fprintf(os.Stderr, "Starting Aula login (%s)...\n", authLevel)
	fmt.Fprintln(os.Stderr)

	// Step 1: authenticate in the browser.
	fmt.Fprintln(os.Stderr, "Step 1: Opening browser for authentication.")
	fmt.Fprintln(os.Stderr, "        Complete the login in your browser.")
	fmt.Fprintln(os.Stderr, "        The browser will end up on Aula's website — that's expected.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "If the browser does not open, visit this URL manually:")
	fmt.Fprintf(os.Stderr, "  %s\n", authorizeURL)
	fmt.Fprintln(os.Stderr)

	if err := openBrowser(authorizeURL); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to open browser: %v\n", err)
		fmt.Fprintln(os.Stderr, "Please open the URL above manually.")
	}

	fmt.Fprint(os.Stderr, "        Press Enter when done...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Fprintln(os.Stderr)

	// Step 2: re-open the same URL to capture the redirect.
	fmt.Fprintln(os.Stderr, "Step 2: Opening browser again to capture the authorization code...")
	fmt.Fprintln(os.Stderr, "        Your browser will land on a page that won't fully load.")
	fmt.Fprintln(os.Stderr)

	if err := openBrowser(authorizeURL); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to open browser: %v\n", err)
		fmt.Fprintln(os.Stderr, "Please open the URL above manually.")
	}

	fmt.Fprintln(os.Stderr, "Copy the FULL URL from your browser's address bar and paste it here:")
	fmt.Fprintln(os.Stderr)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "error: no input received")
		os.Exit(1)
	}
	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		fmt.Fprintln(os.Stderr, "error: empty input")
		os.Exit(1)
	}

	parsed, err := url.Parse(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid URL: %v\n", err)
		os.Exit(1)
	}

	resolved, err := resolveRedirectURL(parsed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	code, err := aulaapi.ExtractCodeFromRedirect(resolved.String(), &state)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to extract authorization code: %v\n", err)
		os.Exit(1)
	}

	completeTokenExchange(authLevel, &endpoints, code, pkce.CodeVerifier, "")
}

func resolveRedirectURL(u *url.URL) (*url.URL, error) {
	if u.Host != "app-redirect.aula.dk" {
		return u, nil
	}
	returnURI := u.Query().Get("returnUri")
	if returnURI == "" {
		return nil, fmt.Errorf("app-redirect URL missing returnUri parameter")
	}
	decoded, err := base64.StdEncoding.DecodeString(returnURI)
	if err != nil {
		return nil, fmt.Errorf("failed to decode returnUri base64: %v", err)
	}
	fmt.Fprintln(os.Stderr, "Decoded redirect URL from app-redirect.aula.dk")
	return url.Parse(string(decoded))
}

// ---------------------------------------------------------------------------
// Shared: token exchange + save
// ---------------------------------------------------------------------------

func completeTokenExchange(authLevel aulaapi.AuthLevel, endpoints *aulaapi.OidcEndpoints, code, codeVerifier, redirectURI string) {
	fmt.Fprintln(os.Stderr, "Authorization code received, exchanging for tokens...")

	httpClient := &http.Client{}
	tokenResponse, err := aulaapi.ExchangeCode(httpClient, endpoints, authLevel, code, codeVerifier, redirectURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: token exchange failed: %v\n", err)
		os.Exit(1)
	}

	loginData := aulaapi.LoginDataFromTokenResponse(tokenResponse, authLevel)
	store := cli.GetTokenStore()
	if err := store.Save(loginData); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to save tokens: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, cli.Bold("Login successful!"))
	fmt.Fprintf(os.Stderr, "  Auth level: %s\n", authLevel)
	if loginData.AccessTokenExpiration != nil {
		fmt.Fprintf(os.Stderr, "  Token expires: %s\n", formatUnixTimestamp(*loginData.AccessTokenExpiration))
	}
	fmt.Fprintf(os.Stderr, "  Tokens saved to: %s\n", store.Dir())
}

// ---------------------------------------------------------------------------
// Logout
// ---------------------------------------------------------------------------

func handleLogout(envOverride string) {
	environment := cli.ResolveEnvironment(envOverride)
	store := cli.GetTokenStore()

	if !store.Exists() {
		fmt.Fprintln(os.Stderr, "No active session found.")
		return
	}

	client, err := aulaapi.NewAulaClientWithConfig(aulaapi.AulaClientConfig{
		Environment: environment,
		APIVersion:  23,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create client: %v\n", err)
		os.Exit(1)
	}

	session, err := aulaapi.NewSession(client, store, aulaapi.DefaultSessionConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create session: %v\n", err)
		os.Exit(1)
	}

	if err := session.Logout(cmd_context()); err != nil {
		fmt.Fprintf(os.Stderr, "warning: logout endpoint call failed: %v\n", err)
		fmt.Fprintln(os.Stderr, "Local tokens have been cleared regardless.")
	} else {
		fmt.Fprintln(os.Stderr, "Logged out successfully. Tokens cleared.")
	}
}

// ---------------------------------------------------------------------------
// Status
// ---------------------------------------------------------------------------

func handleStatus() {
	store := cli.GetTokenStore()

	loginData, err := store.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read tokens: %v\n", err)
		os.Exit(1)
	}
	if loginData == nil {
		fmt.Println("Not logged in.")
		fmt.Println("Run 'aula auth login' to authenticate.")
		return
	}

	fmt.Println(cli.Bold("Logged in"))
	fmt.Printf("  Auth level: %s\n", loginData.AuthLevel)
	hasRefresh := "no"
	if loginData.RefreshToken != nil {
		hasRefresh = "yes"
	}
	fmt.Printf("  Has refresh token: %s\n", hasRefresh)

	if loginData.AccessTokenExpiration != nil {
		exp := *loginData.AccessTokenExpiration
		now := uint64(time.Now().Unix())
		fmt.Printf("  Token expires: %s\n", formatUnixTimestamp(exp))

		if exp > now {
			remaining := exp - now
			mins := remaining / 60
			secs := remaining % 60
			fmt.Printf("  Time remaining: %dm %ds\n", mins, secs)
		} else {
			fmt.Printf("  Status: %s\n", cli.Red("EXPIRED"))
			if loginData.RefreshToken != nil {
				fmt.Println("  Run 'aula auth refresh' to get a new token.")
			} else {
				fmt.Println("  Run 'aula auth login' to re-authenticate.")
			}
		}
	} else {
		fmt.Println("  Token expiry: unknown")
	}

	if loginData.Error != nil {
		fmt.Printf("  Error: %s\n", cli.Red(*loginData.Error))
		if loginData.ErrorDescription != nil {
			fmt.Printf("  Error detail: %s\n", *loginData.ErrorDescription)
		}
	}

	fmt.Printf("  Token store: %s\n", store.Dir())
}

// ---------------------------------------------------------------------------
// Refresh
// ---------------------------------------------------------------------------

func handleRefresh(envOverride string) {
	environment := cli.ResolveEnvironment(envOverride)
	store := cli.GetTokenStore()

	if !store.Exists() {
		fmt.Fprintln(os.Stderr, "No active session found. Run 'aula auth login' first.")
		os.Exit(1)
	}

	client, err := aulaapi.NewAulaClientWithConfig(aulaapi.AulaClientConfig{
		Environment: environment,
		APIVersion:  23,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create client: %v\n", err)
		os.Exit(1)
	}

	session, err := aulaapi.NewSession(client, store, aulaapi.DefaultSessionConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create session: %v\n", err)
		os.Exit(1)
	}

	if err := session.RefreshTokens(cmd_context()); err != nil {
		fmt.Fprintf(os.Stderr, "error: token refresh failed: %v\n", err)
		fmt.Fprintln(os.Stderr, "You may need to run 'aula auth login' again.")
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Token refreshed successfully.")
	if ld := session.LoginData(); ld != nil {
		if ld.AccessTokenExpiration != nil {
			fmt.Fprintf(os.Stderr, "  New expiry: %s\n", formatUnixTimestamp(*ld.AccessTokenExpiration))
		}
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func formatUnixTimestamp(ts uint64) string {
	t := time.Unix(int64(ts), 0).UTC()
	return t.Format("2006-01-02 15:04:05 UTC")
}
