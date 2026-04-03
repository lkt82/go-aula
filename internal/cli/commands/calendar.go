package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewCalendarCmd creates the "calendar" command group.
func NewCalendarCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "calendar",
		Short: "View and manage calendar events",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List calendar events for a date range",
		Run: func(c *cobra.Command, args []string) {
			from, _ := c.Flags().GetString("from")
			to, _ := c.Flags().GetString("to")
			group, _ := c.Flags().GetInt64("group")
			institution, _ := c.Flags().GetInt64("institution")

			if from == "" {
				from = today()
			}
			if to == "" {
				to = daysFromToday(7)
			}

			if c.Flags().Changed("group") {
				calHandleListGroup(group, from, to, *jsonFlag, *envFlag)
			} else {
				var instPtr *int64
				if c.Flags().Changed("institution") {
					instPtr = &institution
				}
				calHandleList(from, to, instPtr, *jsonFlag, *envFlag)
			}
		},
	}
	listCmd.Flags().String("from", "", "Start date (YYYY-MM-DD). Defaults to today")
	listCmd.Flags().String("to", "", "End date (YYYY-MM-DD). Defaults to 7 days from start")
	listCmd.Flags().Int64("group", 0, "Filter by group ID")
	listCmd.Flags().Int64("institution", 0, "Filter by institution profile ID")
	cmd.AddCommand(listCmd)

	// show
	showCmd := &cobra.Command{
		Use:   "show <event-id>",
		Short: "Show details for a single event",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			eventID := mustParseInt64(args[0])
			calHandleShow(eventID, *jsonFlag, *envFlag)
		},
	}
	cmd.AddCommand(showCmd)

	// today
	cmd.AddCommand(&cobra.Command{
		Use:   "today",
		Short: "Show today's events",
		Run: func(c *cobra.Command, args []string) {
			calHandleList(today(), daysFromToday(1), nil, *jsonFlag, *envFlag)
		},
	})

	// week
	cmd.AddCommand(&cobra.Command{
		Use:   "week",
		Short: "Show this week's events",
		Run: func(c *cobra.Command, args []string) {
			calHandleList(today(), daysFromToday(7), nil, *jsonFlag, *envFlag)
		},
	})

	// respond
	respondCmd := &cobra.Command{
		Use:   "respond <event-id>",
		Short: "Respond to an event invitation",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			eventID := mustParseInt64(args[0])
			accept, _ := c.Flags().GetBool("accept")
			decline, _ := c.Flags().GetBool("decline")
			tentative, _ := c.Flags().GetBool("tentative")
			profile, _ := c.Flags().GetInt64("profile")
			var profilePtr *int64
			if c.Flags().Changed("profile") {
				profilePtr = &profile
			}
			calHandleRespond(eventID, accept, decline, tentative, profilePtr, *jsonFlag, *envFlag)
		},
	}
	respondCmd.Flags().Bool("accept", false, "Accept the invitation")
	respondCmd.Flags().Bool("decline", false, "Decline the invitation")
	respondCmd.Flags().Bool("tentative", false, "Respond tentatively")
	respondCmd.Flags().Int64("profile", 0, "Institution profile ID (required for response)")
	cmd.AddCommand(respondCmd)

	// birthdays
	birthdaysCmd := &cobra.Command{
		Use:   "birthdays",
		Short: "Show birthdays for a group or institution",
		Run: func(c *cobra.Command, args []string) {
			from, _ := c.Flags().GetString("from")
			to, _ := c.Flags().GetString("to")
			group, _ := c.Flags().GetInt64("group")
			institution, _ := c.Flags().GetInt64("institution")
			if from == "" {
				from = today()
			}
			if to == "" {
				to = daysFromToday(30)
			}
			var groupPtr, instPtr *int64
			if c.Flags().Changed("group") {
				groupPtr = &group
			}
			if c.Flags().Changed("institution") {
				instPtr = &institution
			}
			calHandleBirthdays(groupPtr, instPtr, from, to, *jsonFlag, *envFlag)
		},
	}
	birthdaysCmd.Flags().Int64("group", 0, "Group ID")
	birthdaysCmd.Flags().Int64("institution", 0, "Institution ID")
	birthdaysCmd.Flags().String("from", "", "Start date (YYYY-MM-DD)")
	birthdaysCmd.Flags().String("to", "", "End date (YYYY-MM-DD)")
	cmd.AddCommand(birthdaysCmd)

	return cmd
}

func today() string {
	return time.Now().Format("2006-01-02")
}

func daysFromToday(days int) string {
	return time.Now().AddDate(0, 0, days).Format("2006-01-02")
}

func calHandleList(start, end string, instProfileID *int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	var instProfileIDs []int64
	if instProfileID != nil {
		instProfileIDs = []int64{*instProfileID}
	} else {
		if err := session.EnsureContextInitialized(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to initialize session: %v\n", err)
			os.Exit(1)
		}
		instProfileIDs = session.AllInstitutionProfileIDs()
	}

	params := &models.GetEventsParameters{
		InstProfileIDs: instProfileIDs,
		Start:          &start,
		End:            &end,
	}

	events, err := services.GetEvents(ctx, session, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(events)
	} else {
		printEventList(events)
	}
}

func calHandleListGroup(groupID int64, start, end string, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	events, err := services.GetEventForGroup(ctx, session, groupID, &start, &end)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(events)
	} else {
		printEventList(events)
	}
}

func printEventList(events []models.EventSimpleDto) {
	if len(events) == 0 {
		fmt.Println("No events found.")
		return
	}

	table := cli.NewTable([]cli.Column{
		{Header: "ID", Width: 12},
		{Header: "DATE", Width: 12},
		{Header: "TIME", Width: 7},
		{Header: "TITLE", Width: 30},
		{Header: "TYPE", Width: 14},
		{Header: "RESPONSE", Width: 10},
	})
	table.PrintHeader()

	for _, event := range events {
		id := "-"
		if event.ID != nil {
			id = fmt.Sprintf("%d", *event.ID)
		}

		date, timeStr := cli.SplitDatetime(event.StartDateTime)
		allDay := event.AllDay != nil && *event.AllDay
		timeDisplay := timeStr
		if allDay {
			timeDisplay = "all-day"
		}

		title := "(untitled)"
		if event.Title != nil {
			title = aulaapi.ExpandTitle(*event.Title)
		}
		eventType := ""
		if event.EventType != nil {
			eventType = *event.EventType
		}
		response := ""
		if event.ResponseStatus != nil {
			response = *event.ResponseStatus
		}

		table.PrintRow([]string{id, date, timeDisplay, title, eventType, response})
	}

	fmt.Printf("\n%d event(s) total.\n", len(events))
}

func calHandleShow(eventID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	detail, err := services.GetEventDetail(ctx, session, eventID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(detail)
	} else {
		printEventDetail(&detail)
	}
}

func printEventDetail(detail *models.EventDetailsDto) {
	title := "(untitled)"
	if detail.Title != nil {
		title = *detail.Title
	}
	eventType := "unknown"
	if detail.EventType != nil {
		eventType = *detail.EventType
	}

	fmt.Println(cli.Bold(fmt.Sprintf("Event: %s", title)))
	fmt.Printf("  Type: %s\n", eventType)

	if detail.ID != nil {
		fmt.Printf("  ID: %d\n", *detail.ID)
	}
	if detail.StartDateTime != nil {
		fmt.Printf("  Start: %s", cli.FormatDatetime(*detail.StartDateTime))
	}
	if detail.EndDateTime != nil {
		fmt.Printf("  End: %s", cli.FormatDatetime(*detail.EndDateTime))
	}
	if detail.AllDay != nil && *detail.AllDay {
		fmt.Print("  (all day)")
	}
	fmt.Println()

	if detail.PrimaryResourceText != nil {
		fmt.Printf("  Location: %s\n", *detail.PrimaryResourceText)
	} else if detail.PrimaryResource != nil && detail.PrimaryResource.Name != nil {
		fmt.Printf("  Resource: %s\n", *detail.PrimaryResource.Name)
	}

	if detail.InstitutionCode != nil {
		fmt.Printf("  Institution: %s\n", *detail.InstitutionCode)
	}

	if detail.ResponseStatus != nil {
		fmt.Printf("  Your response: %s\n", *detail.ResponseStatus)
	}

	if detail.ResponseRequired != nil && *detail.ResponseRequired {
		fmt.Print("  [Response required]")
		if detail.ResponseDeadline != nil {
			fmt.Printf("  Deadline: %s", cli.FormatDatetime(*detail.ResponseDeadline))
		}
		fmt.Println()
	}

	if detail.Creator != nil && detail.Creator.Name != nil {
		fmt.Printf("  Created by: %s\n", *detail.Creator.Name)
	}

	if detail.Description != nil && detail.Description.HTML != nil {
		plain := cli.StripHTMLTags(*detail.Description.HTML)
		if trimmed := trimString(plain); trimmed != "" {
			fmt.Println()
			fmt.Printf("%s:\n", cli.Bold("Description"))
			fmt.Println(trimmed)
		}
	}

	if len(detail.InvitedGroups) > 0 {
		fmt.Println()
		fmt.Printf("%s:\n", cli.Bold("Invited groups"))
		for _, g := range detail.InvitedGroups {
			name := "(unnamed)"
			if g.Name != nil {
				name = *g.Name
			}
			fmt.Printf("  - %s\n", name)
		}
	}

	if len(detail.Invitees) > 0 {
		fmt.Println()
		fmt.Printf("%s (%d):\n", cli.Bold("Invitees"), len(detail.Invitees))
		for _, inv := range detail.Invitees {
			name := "(unknown)"
			if inv.InstProfile != nil && inv.InstProfile.FullName != nil {
				name = *inv.InstProfile.FullName
			}
			response := "N/A"
			if inv.ResponseType != nil {
				response = *inv.ResponseType
			}
			fmt.Printf("  - %s [%s]\n", name, response)
		}
	}

	if len(detail.Attachments) > 0 {
		fmt.Println()
		fmt.Printf("%s:\n", cli.Bold("Attachments"))
		for _, att := range detail.Attachments {
			name := "(unnamed)"
			if att.File != nil && att.File.Name != nil {
				name = *att.File.Name
			}
			fmt.Printf("  - %s\n", name)
		}
	}
}

func calHandleRespond(eventID int64, accept, decline, tentative bool, profileID *int64, jsonOut bool, envOverride string) {
	var responseType string
	switch {
	case accept:
		responseType = "Accepted"
	case decline:
		responseType = "Declined"
	case tentative:
		responseType = "Tentative"
	default:
		fmt.Fprintln(os.Stderr, "error: specify one of --accept, --decline, or --tentative")
		os.Exit(1)
	}

	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	args := &models.RespondSimpleEventRequest{
		EventID:                  &eventID,
		InstitutionProfileID:     profileID,
		InvitedInstProfileID:     profileID,
		ResponseType:             &responseType,
	}

	result, err := services.RespondSimpleEvent(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Responded to event %d: %s\n", eventID, responseType)
	}
}

func calHandleBirthdays(groupID, institutionID *int64, start, end string, jsonOut bool, envOverride string) {
	if groupID == nil && institutionID == nil {
		fmt.Fprintln(os.Stderr, "error: specify --group <id> or --institution <id>")
		os.Exit(1)
	}

	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	var birthdays []models.BirthdayEventDto
	var err error
	if groupID != nil {
		birthdays, err = services.GetBirthdaysForGroup(ctx, session, *groupID, start, end)
	} else {
		birthdays, err = services.GetBirthdaysForInstitution(ctx, session, *institutionID, start, end)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if jsonOut {
		cli.PrintJSON(birthdays)
	} else if len(birthdays) == 0 {
		fmt.Println("No birthdays found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "DATE", Width: 12},
			{Header: "NAME", Width: 25},
			{Header: "GROUP", Width: 20},
		})
		table.PrintHeader()
		for _, bday := range birthdays {
			date := ""
			if bday.Birthday != nil {
				date = *bday.Birthday
			}
			name := "(unknown)"
			if bday.Name != nil {
				name = *bday.Name
			}
			group := ""
			if bday.MainGroupName != nil {
				group = *bday.MainGroupName
			}
			table.PrintRow([]string{date, name, group})
		}
		fmt.Printf("\n%d birthday(s) total.\n", len(birthdays))
	}
}
