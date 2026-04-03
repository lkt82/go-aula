package cli

import (
	"fmt"
	"os"

	"github.com/lkt82/go-aula/internal/aulaapi"
)

// ResolveEnvironment maps a CLI environment string to an Environment value.
// Falls back to production if the string is empty or unrecognized.
func ResolveEnvironment(envStr string) aulaapi.Environment {
	if envStr == "" {
		return aulaapi.EnvProduction
	}
	env, err := aulaapi.ParseEnvironment(envStr)
	if err != nil {
		return aulaapi.EnvProduction
	}
	return env
}

// GetTokenStore returns the default token store, with a fallback.
func GetTokenStore() *aulaapi.TokenStore {
	store, err := aulaapi.DefaultTokenStore()
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning: could not determine data directory, using ./aula-data")
		return aulaapi.NewTokenStore("./aula-data")
	}
	return store
}

// BuildSession builds an authenticated session, exiting on failure.
func BuildSession(envOverride string) *aulaapi.Session {
	environment := ResolveEnvironment(envOverride)
	store := GetTokenStore()

	if !store.Exists() {
		fmt.Fprintln(os.Stderr, "Not logged in. Run 'aula auth login' first.")
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

	return session
}
