package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewProfileCmd creates the "profile" command group.
func NewProfileCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "View and manage user profiles",
	}

	// me
	cmd.AddCommand(&cobra.Command{
		Use:   "me",
		Short: "Show current user's profile(s) from login data",
		Run: func(c *cobra.Command, args []string) {
			profileHandleMe(*jsonFlag, *envFlag)
		},
	})

	// master-data
	cmd.AddCommand(&cobra.Command{
		Use:   "master-data",
		Short: "Show profile master data (contact details)",
		Run: func(c *cobra.Command, args []string) {
			profileHandleMasterData(*jsonFlag, *envFlag)
		},
	})

	return cmd
}

func profileHandleMe(jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	resp, err := services.GetProfilesByLogin(ctx, session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(resp)
	} else if len(resp.Profiles) == 0 {
		fmt.Println("No profiles found.")
	} else {
		for i, profile := range resp.Profiles {
			if i > 0 {
				fmt.Println(cli.Dim("--------------------------------------------------"))
			}
			name := "(unknown)"
			if profile.DisplayName != nil {
				name = *profile.DisplayName
			}
			role := "(unknown)"
			if profile.PortalRole != nil {
				role = *profile.PortalRole
			}

			fmt.Println(cli.Bold(fmt.Sprintf("Profile #%d", i+1)))
			fmt.Printf("  Name: %s\n", name)
			fmt.Printf("  Role: %s\n", role)
			if profile.ProfileID != nil {
				fmt.Printf("  Profile ID: %d\n", *profile.ProfileID)
			}

			for _, ip := range profile.InstitutionProfiles {
				fmt.Println()
				instName := "(unknown)"
				if ip.InstitutionName != nil {
					instName = *ip.InstitutionName
				}
				fmt.Printf("  %s\n", cli.Bold(fmt.Sprintf("Institution: %s", instName)))
				fmt.Printf("    Institution profile ID: %d\n", ip.ID)
				if ip.InstitutionCode != nil {
					fmt.Printf("    Institution code: %s\n", *ip.InstitutionCode)
				}
				if ip.MunicipalityName != nil {
					fmt.Printf("    Municipality: %s\n", *ip.MunicipalityName)
				}
				if ip.Email != nil {
					fmt.Printf("    Email: %s\n", *ip.Email)
				}
				if ip.MobilePhoneNumber != nil {
					fmt.Printf("    Mobile: %s\n", *ip.MobilePhoneNumber)
				}
				if ip.Address != nil {
					street := ""
					if ip.Address.Street != nil {
						street = *ip.Address.Street
					}
					postal := ""
					if ip.Address.PostalCode != nil {
						postal = *ip.Address.PostalCode
					}
					district := ""
					if ip.Address.PostalDistrict != nil {
						district = *ip.Address.PostalDistrict
					}
					if street != "" {
						fmt.Printf("    Address: %s, %s %s\n", street, postal, district)
					}
				}
			}

			// Children
			var childNames []string
			for _, c := range profile.Children {
				if c.Name != nil {
					childNames = append(childNames, *c.Name)
				}
			}
			if len(childNames) > 0 {
				fmt.Println()
				fmt.Printf("  Children (%d): %s\n", len(childNames), joinStrings(childNames, ", "))
			}
		}
	}
}

func profileHandleMasterData(jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	profile, err := services.GetProfileMasterData(ctx, session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(profile)
	} else {
		first := ""
		if profile.FirstName != nil {
			first = *profile.FirstName
		}
		last := ""
		if profile.LastName != nil {
			last = *profile.LastName
		}
		role := "(unknown)"
		if profile.PortalRole != nil {
			role = *profile.PortalRole
		}

		fmt.Println(cli.Bold("Master Data"))
		fmt.Printf("  Name: %s %s\n", first, last)
		fmt.Printf("  Role: %s\n", role)
		if profile.ExternalEmail != nil {
			fmt.Printf("  Email: %s\n", *profile.ExternalEmail)
		}
		if profile.Phonenumber != nil {
			fmt.Printf("  Phone: %s\n", *profile.Phonenumber)
		}
		if profile.MobilePhonenumber != nil {
			fmt.Printf("  Mobile: %s\n", *profile.MobilePhonenumber)
		}
		if profile.WorkPhonenumber != nil {
			fmt.Printf("  Work phone: %s\n", *profile.WorkPhonenumber)
		}
		if profile.HomePhonenumber != nil {
			fmt.Printf("  Home phone: %s\n", *profile.HomePhonenumber)
		}
	}
}
