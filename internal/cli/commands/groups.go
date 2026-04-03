package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewGroupsCmd creates the "groups" command group.
func NewGroupsCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "View and manage groups",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List groups for an institution profile (context)",
		Run: func(c *cobra.Command, args []string) {
			instProfile, _ := c.Flags().GetInt64("inst-profile")
			groupsHandleList(instProfile, *jsonFlag, *envFlag)
		},
	}
	listCmd.Flags().Int64("inst-profile", 0, "Institution profile ID to list groups for")
	_ = listCmd.MarkFlagRequired("inst-profile")
	cmd.AddCommand(listCmd)

	// show
	cmd.AddCommand(&cobra.Command{
		Use:   "show <group-id>",
		Short: "Show group details",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			groupID := mustParseInt64(args[0])
			groupsHandleShow(groupID, *jsonFlag, *envFlag)
		},
	})

	// members
	cmd.AddCommand(&cobra.Command{
		Use:   "members <group-id>",
		Short: "List members of a group",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			groupID := mustParseInt64(args[0])
			groupsHandleMembers(groupID, *jsonFlag, *envFlag)
		},
	})

	return cmd
}

func groupsHandleList(profileID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	groupList, err := services.GetGroupByContext(ctx, session, profileID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(groupList)
	} else if len(groupList) == 0 {
		fmt.Println("No groups found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "ID", Width: 8},
			{Header: "NAME", Width: 30},
			{Header: "DEFAULT", Width: 10},
		})
		table.PrintHeader()
		for _, g := range groupList {
			id := ""
			if g.ID != nil {
				id = fmt.Sprintf("%d", *g.ID)
			}
			name := "(unnamed)"
			if g.Name != nil {
				name = *g.Name
			}
			def := "no"
			if g.ShowAsDefault {
				def = "yes"
			}
			table.PrintRow([]string{id, cli.Truncate(name, 30), def})
		}
	}
}

func groupsHandleShow(groupID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	group, err := services.GetGroup(ctx, session, groupID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(group)
	} else {
		name := "(unnamed)"
		if group.Name != nil {
			name = *group.Name
		}
		fmt.Println(cli.Bold(fmt.Sprintf("Group: %s", name)))
		if group.Description != nil && *group.Description != "" {
			fmt.Printf("  Description: %s\n", *group.Description)
		}
		if group.GroupType != nil {
			fmt.Printf("  Type: %s\n", *group.GroupType)
		}
		if group.Status != nil {
			fmt.Printf("  Status: %s\n", *group.Status)
		}
		if group.Access != nil {
			fmt.Printf("  Access: %s\n", *group.Access)
		}
		if group.Role != nil {
			fmt.Printf("  Your role: %s\n", *group.Role)
		}
		if group.InstitutionCode != nil {
			fmt.Printf("  Institution: %s\n", *group.InstitutionCode)
		}
		if group.DashboardEnabled {
			fmt.Println("  Dashboard: enabled")
		}
	}
}

func groupsHandleMembers(groupID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	members, err := services.GetMembershipsLight(ctx, session, groupID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(members)
	} else if len(members) == 0 {
		fmt.Println("No members found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "ID", Width: 8},
			{Header: "ROLE", Width: 15},
			{Header: "INST. ROLE", Width: 15},
		})
		table.PrintHeader()
		for _, m := range members {
			id := ""
			if m.ID != nil {
				id = fmt.Sprintf("%d", *m.ID)
			}
			role := ""
			if m.GroupRole != nil {
				role = *m.GroupRole
			}
			instRole := ""
			if m.InstitutionRole != nil {
				instRole = *m.InstitutionRole
			}
			table.PrintRow([]string{id, role, instRole})
		}
		fmt.Printf("\n%d member(s)\n", len(members))
	}
}
