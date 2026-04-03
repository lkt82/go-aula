package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewNotificationsCmd creates the "notifications" command group.
func NewNotificationsCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notifications",
		Short: "View and manage notifications",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List recent notifications",
		Run: func(c *cobra.Command, args []string) {
			limit, _ := c.Flags().GetInt32("limit")
			all, _ := c.Flags().GetBool("all")
			lim := int(limit)
			if all {
				lim = 1<<31 - 1
			}
			notifHandleList(lim, *jsonFlag, *envFlag)
		},
	}
	listCmd.Flags().Int32P("limit", "n", 20, "Maximum number of notifications to show")
	listCmd.Flags().Bool("all", false, "Show all notifications (ignore limit)")
	cmd.AddCommand(listCmd)

	// delete-all
	cmd.AddCommand(&cobra.Command{
		Use:   "delete-all",
		Short: "Delete all notifications",
		Run: func(c *cobra.Command, args []string) {
			notifHandleDeleteAll(*jsonFlag, *envFlag)
		},
	})

	// delete-child
	cmd.AddCommand(&cobra.Command{
		Use:   "delete-child <child-id>",
		Short: "Delete notifications for a specific child",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			childID := mustParseInt64(args[0])
			notifHandleDeleteChild(childID, *jsonFlag, *envFlag)
		},
	})

	return cmd
}

func notifHandleList(limit int, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	if err := session.EnsureContextInitialized(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to initialize session: %v\n", err)
		os.Exit(1)
	}
	childrenIDs := session.ChildrenInstProfileIDs()
	institutionCodes := session.ChildrenInstitutionCodes()

	items, err := services.GetNotifications(ctx, session, childrenIDs, institutionCodes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(items)
	} else if len(items) == 0 {
		fmt.Println("No notifications.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "TYPE", Width: 20},
			{Header: "AREA", Width: 15},
			{Header: "TITLE", Width: 40},
		})
		table.PrintHeader()
		for i, item := range items {
			if i >= limit {
				break
			}
			eventType := ""
			if item.NotificationEventType != nil {
				eventType = *item.NotificationEventType
			}
			area := ""
			if item.NotificationArea != nil {
				area = *item.NotificationArea
			}
			title := "(no title)"
			if item.Title != nil {
				title = *item.Title
			}
			table.PrintRow([]string{
				cli.Truncate(eventType, 20),
				cli.Truncate(area, 15),
				cli.Truncate(title, 40),
			})
		}
		if len(items) > limit {
			fmt.Fprintf(os.Stderr, "\n(showing %d of %d notifications)\n", limit, len(items))
		}
	}
}

func notifHandleDeleteAll(jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	result, err := services.DeleteNotifications(ctx, session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Println("All notifications deleted.")
	}
}

func notifHandleDeleteChild(childID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	result, err := services.DeleteNotificationForChild(ctx, session, childID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Notifications for child %d deleted.\n", childID)
	}
}
