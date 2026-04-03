package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewConfigCmd creates the "config" command group.
func NewConfigCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "View and manage CLI configuration",
	}

	// show
	cmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "Show current configuration (merged from file and defaults)",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("config show: not yet implemented")
		},
	})

	// path
	cmd.AddCommand(&cobra.Command{
		Use:   "path",
		Short: "Show the configuration file path",
		Run: func(c *cobra.Command, args []string) {
			path := cli.ConfigPath()
			if path == "" {
				fmt.Println("Could not determine config directory")
			} else {
				fmt.Println(path)
			}
		},
	})

	// set
	cmd.AddCommand(&cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration value",
		Args:  cobra.ExactArgs(2),
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("config set: not yet implemented")
		},
	})

	// init
	cmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Initialize a default configuration file",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("config init: not yet implemented")
		},
	})

	// policy
	policyCmd := &cobra.Command{
		Use:   "policy",
		Short: "Show policy links (data policy, terms of use, etc.)",
		Run: func(c *cobra.Command, args []string) {
			jsonOut := *jsonFlag
			session := cli.BuildSession(*envFlag)
			ctx := context.Background()

			entries, err := services.GetPolicyLinks(ctx, session)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			if jsonOut {
				cli.PrintJSON(entries)
			} else if len(entries) == 0 {
				fmt.Println("No policy documents found.")
			} else {
				for _, entry := range entries {
					title := "(untitled)"
					if entry.CommonFile != nil && entry.CommonFile.Title != nil {
						title = *entry.CommonFile.Title
					}
					inst := ""
					if entry.Institution != nil && entry.Institution.InstitutionName != nil {
						inst = *entry.Institution.InstitutionName
					}
					url := ""
					if entry.CommonFile != nil && entry.CommonFile.File != nil && entry.CommonFile.File.File != nil && entry.CommonFile.File.File.URL != nil {
						url = *entry.CommonFile.File.File.URL
					}
					fmt.Println(cli.Bold(title))
					if inst != "" {
						fmt.Printf("  Institution: %s\n", inst)
					}
					if url != "" {
						fmt.Printf("  URL: %s\n", url)
					}
					fmt.Println()
				}
			}
		},
	}
	cmd.AddCommand(policyCmd)

	// privacy
	privacyCmd := &cobra.Command{
		Use:   "privacy",
		Short: "Show the privacy/data policy content",
		Run: func(c *cobra.Command, args []string) {
			jsonOut := *jsonFlag
			session := cli.BuildSession(*envFlag)
			ctx := context.Background()

			policy, err := services.GetPrivacyPolicy(ctx, session)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			if jsonOut {
				cli.PrintJSON(policy)
			} else {
				if policy.Version != nil {
					fmt.Printf("Version: %s\n\n", *policy.Version)
				}
				if len(policy.Content) > 0 {
					var wrapper struct {
						HTML string `json:"html"`
					}
					if err := json.Unmarshal(policy.Content, &wrapper); err == nil && wrapper.HTML != "" {
						fmt.Println(cli.StripHTMLTags(wrapper.HTML))
					} else {
						fmt.Println(string(policy.Content))
					}
				} else {
					fmt.Println("No policy content returned.")
				}
			}
		},
	}
	cmd.AddCommand(privacyCmd)

	return cmd
}
