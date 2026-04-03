// aula-cli -- command-line interface for the Aula school platform.
//
// Built on top of the aulaapi library. Provides subcommands for each
// major domain: auth, messages, calendar, presence, posts, gallery, documents,
// notifications, search, groups, profile, and config.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/cli"
	"github.com/lkt82/go-aula/internal/cli/commands"
)

const version = "0.1.0"

func main() {
	var (
		jsonOut bool
		envStr  string
		verbose bool
		profile string
		info    bool
	)

	rootCmd := &cobra.Command{
		Use:   "aula",
		Short: "CLI tool for interacting with the Aula school platform",
		Long: `Aula is Denmark's school communication platform by Netcompany A/S.
This tool provides command-line access to messages, calendar, presence,
posts, gallery, documents, notifications, search, and more.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cfg := cli.LoadConfig()
			if verbose || cfg.Verbose {
				fmt.Fprintf(os.Stderr, "[verbose] config loaded from %s\n", cli.ConfigPath())
				if envStr != "" {
					fmt.Fprintf(os.Stderr, "[verbose] env override: %s\n", envStr)
				}
				if jsonOut {
					fmt.Fprintln(os.Stderr, "[verbose] JSON output enabled")
				}
			}

			// Resolve environment: CLI flag > config file > default (production).
			if envStr == "" && cfg.DefaultEnvironment != "" {
				envStr = cfg.DefaultEnvironment
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if info {
				fmt.Printf("aula-cli v%s\n", version)
				return
			}
			cmd.Help()
		},
	}

	// Global flags.
	rootCmd.PersistentFlags().BoolVar(&jsonOut, "json", false, "Output results as JSON instead of human-readable text")
	rootCmd.PersistentFlags().StringVar(&envStr, "env", "", "Aula environment (production, preprod, hotfix, test1, test3, dev1, dev3, dev11)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output (debug logging)")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "", "Institution profile selector")
	rootCmd.Flags().BoolVar(&info, "info", false, "Print version information and exit")

	// Register all subcommands.
	rootCmd.AddCommand(commands.NewAuthCmd(&envStr))
	rootCmd.AddCommand(commands.NewMessagesCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewCalendarCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewPresenceCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewPostsCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewGalleryCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewDocumentsCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewNotificationsCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewSearchCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewGroupsCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewProfileCmd(&jsonOut, &envStr))
	rootCmd.AddCommand(commands.NewConfigCmd(&jsonOut, &envStr))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
