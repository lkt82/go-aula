package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewPresenceCmd creates the "presence" command group.
func NewPresenceCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "presence",
		Short: "View and manage child presence (attendance)",
	}

	// status
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show current presence status for children",
		Run: func(c *cobra.Command, args []string) {
			children, _ := c.Flags().GetInt64Slice("children")
			presHandleStatus(children, *jsonFlag, *envFlag)
		},
	}
	statusCmd.Flags().Int64Slice("children", nil, "Institution profile IDs (comma-separated)")
	cmd.AddCommand(statusCmd)

	// registrations
	regCmd := &cobra.Command{
		Use:   "registrations",
		Short: "Show presence registrations for a date",
		Run: func(c *cobra.Command, args []string) {
			children, _ := c.Flags().GetInt64Slice("children")
			presHandleRegistrations(children, *jsonFlag, *envFlag)
		},
	}
	regCmd.Flags().Int64Slice("children", nil, "Institution profile IDs (comma-separated)")
	regCmd.Flags().String("date", "", "Date to query (YYYY-MM-DD)")
	cmd.AddCommand(regCmd)

	// schedule
	schedCmd := &cobra.Command{
		Use:   "schedule",
		Short: "Show weekly presence schedule",
		Run: func(c *cobra.Command, args []string) {
			children, _ := c.Flags().GetInt64Slice("children")
			from, _ := c.Flags().GetString("from")
			to, _ := c.Flags().GetString("to")
			presHandleSchedule(children, from, to, *jsonFlag, *envFlag)
		},
	}
	schedCmd.Flags().Int64Slice("children", nil, "Institution profile IDs (comma-separated)")
	schedCmd.Flags().String("from", "", "Start date (YYYY-MM-DD)")
	schedCmd.Flags().String("to", "", "End date (YYYY-MM-DD)")
	cmd.AddCommand(schedCmd)

	// report-status
	reportCmd := &cobra.Command{
		Use:   "report-status",
		Short: "Report a child as sick or absent by institution profile ID",
		Run: func(c *cobra.Command, args []string) {
			children, _ := c.Flags().GetInt64Slice("children")
			status, _ := c.Flags().GetInt32("status")
			presHandleReportStatus(children, int(status), *jsonFlag, *envFlag)
		},
	}
	reportCmd.Flags().Int64Slice("children", nil, "Child institution profile IDs (comma-separated)")
	reportCmd.Flags().Int32("status", 0, "Status code: 0=NotPresent, 1=Sick, 2=ReportedAbsence, 3=Present")
	_ = reportCmd.MarkFlagRequired("children")
	_ = reportCmd.MarkFlagRequired("status")
	cmd.AddCommand(reportCmd)

	return cmd
}

func resolveChildrenIDs(ctx context.Context, session interface {
	EnsureContextInitialized(context.Context) error
	ChildrenInstProfileIDs() []int64
}, children []int64) []int64 {
	if len(children) > 0 {
		return children
	}
	if err := session.EnsureContextInitialized(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to initialize session: %v\n", err)
		os.Exit(1)
	}
	ids := session.ChildrenInstProfileIDs()
	if len(ids) == 0 {
		fmt.Fprintln(os.Stderr, "error: no children found in profile; specify --children explicitly")
		os.Exit(1)
	}
	return ids
}

func presHandleStatus(children []int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()
	children = resolveChildrenIDs(ctx, session, children)

	states, err := services.GetChildrensState(ctx, session, children)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(states)
	} else if len(states) == 0 {
		fmt.Println("No presence status found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "PROFILE ID", Width: 12},
			{Header: "STATUS", Width: 15},
			{Header: "NAME", Width: 20},
		})
		table.PrintHeader()
		for _, s := range states {
			statusRaw := "(unknown)"
			if s.State != nil {
				statusRaw = s.State.String()
			}
			statusDisplay := cli.ColorPresenceStatus(statusRaw)
			name := "(unknown)"
			if s.UniStudent != nil && s.UniStudent.Name != nil {
				name = *s.UniStudent.Name
			}
			profileID := fmt.Sprintf("%d", s.UniStudentID)
			table.PrintColoredRow(
				[]string{profileID, statusRaw, name},
				[]string{profileID, statusDisplay, name},
			)
		}
	}
}

func presHandleRegistrations(children []int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()
	children = resolveChildrenIDs(ctx, session, children)

	regs, err := services.GetDailyOverview(ctx, session, children)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(regs)
	} else if len(regs) == 0 {
		fmt.Println("No presence registrations found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "CHILD", Width: 20},
			{Header: "STATUS", Width: 15},
			{Header: "CHECK-IN", Width: 12},
			{Header: "CHECK-OUT", Width: 12},
			{Header: "LOCATION", Width: 15},
		})
		table.PrintHeader()
		for _, r := range regs {
			statusRaw := ""
			if r.Status != nil {
				statusRaw = presenceStatusName(*r.Status)
			}
			statusDisplay := cli.ColorPresenceStatus(statusRaw)
			checkin := ""
			if r.CheckInTime != nil {
				checkin = cli.ExtractTime(*r.CheckInTime)
			}
			checkout := ""
			if r.CheckOutTime != nil {
				checkout = cli.ExtractTime(*r.CheckOutTime)
			}
			childName := "(unknown)"
			if r.InstitutionProfile != nil && r.InstitutionProfile.Name != nil {
				childName = *r.InstitutionProfile.Name
			}
			location := ""
			if r.Location != nil && r.Location.Name != nil {
				location = cli.Truncate(*r.Location.Name, 15)
			}
			table.PrintColoredRow(
				[]string{childName, statusRaw, checkin, checkout, location},
				[]string{childName, statusDisplay, checkin, checkout, location},
			)
		}
	}
}

func presHandleSchedule(children []int64, from, to string, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()
	children = resolveChildrenIDs(ctx, session, children)

	now := time.Now()
	if from == "" {
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		monday := now.AddDate(0, 0, -(weekday - 1))
		from = monday.Format("2006-01-02")
	}
	if to == "" {
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		sunday := now.AddDate(0, 0, 7-weekday)
		to = sunday.Format("2006-01-02")
	}

	args := &models.PresenceSchedulesRequest{
		FilterInstitutionProfileIDs: children,
		FromDate:                    &from,
		ToDate:                      &to,
	}

	schedule, err := services.GetPresenceSchedules(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(schedule)
	} else {
		presDisplaySchedule(schedule)
	}
}

type scheduleWeekTemplate struct {
	InstitutionProfile struct {
		Name            string `json:"name"`
		InstitutionName string `json:"institutionName"`
	} `json:"institutionProfile"`
	DayTemplates []struct {
		DayOfWeek  int     `json:"dayOfWeek"`
		EntryTime  *string `json:"entryTime,omitempty"`
		ExitTime   *string `json:"exitTime,omitempty"`
		IsOnVacation bool  `json:"isOnVacation"`
		ExitWith   *string `json:"exitWith,omitempty"`
		Comment    *string `json:"comment,omitempty"`
	} `json:"dayTemplates"`
}

var dayNames = []string{"", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func presDisplaySchedule(raw json.RawMessage) {
	var schedule struct {
		CurrentDate           *string            `json:"currentDate,omitempty"`
		PresenceWeekTemplates []json.RawMessage  `json:"presenceWeekTemplates,omitempty"`
	}
	if err := json.Unmarshal(raw, &schedule); err != nil {
		fmt.Println("(failed to parse schedule)")
		return
	}
	if schedule.CurrentDate != nil {
		fmt.Printf("Week of %s\n\n", *schedule.CurrentDate)
	}
	for _, raw := range schedule.PresenceWeekTemplates {
		var tmpl scheduleWeekTemplate
		if err := json.Unmarshal(raw, &tmpl); err != nil {
			continue
		}
		fmt.Println(cli.Bold(fmt.Sprintf("%s (%s)", tmpl.InstitutionProfile.Name, tmpl.InstitutionProfile.InstitutionName)))
		table := cli.NewTable([]cli.Column{
			{Header: "DAY", Width: 5},
			{Header: "ENTRY", Width: 8},
			{Header: "EXIT", Width: 8},
			{Header: "NOTE", Width: 30},
		})
		table.PrintHeader()
		for _, d := range tmpl.DayTemplates {
			day := ""
			if d.DayOfWeek >= 1 && d.DayOfWeek <= 7 {
				day = dayNames[d.DayOfWeek]
			}
			entry := "-"
			exit := "-"
			note := ""
			if d.IsOnVacation {
				note = "Vacation"
			} else {
				if d.EntryTime != nil && *d.EntryTime != "" {
					entry = *d.EntryTime
				}
				if d.ExitTime != nil && *d.ExitTime != "" {
					exit = *d.ExitTime
				}
				if d.ExitWith != nil && *d.ExitWith != "" {
					note = "exit with: " + *d.ExitWith
				}
				if d.Comment != nil && *d.Comment != "" {
					if note != "" {
						note += " "
					}
					note += "[" + *d.Comment + "]"
				}
			}
			table.PrintRow([]string{day, entry, exit, note})
		}
		fmt.Println()
	}
	if len(schedule.PresenceWeekTemplates) == 0 {
		fmt.Println("No schedule data found.")
	}
}

func presHandleReportStatus(children []int64, status int, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	args := &models.UpdateStatusByInstitutionProfileIds{
		InstitutionProfileIDs: children,
		Status:                status,
	}

	result, err := services.UpdateStatusByInstitutionProfileIDs(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Status updated to %s for %d profile(s).\n",
			cli.ColorPresenceStatus(presenceStatusName(status)), len(children))
	}
}

func presenceStatusName(status int) string {
	switch status {
	case 0:
		return "NotPresent"
	case 1:
		return "Sick"
	case 2:
		return "ReportedAbsence"
	case 3:
		return "Present"
	default:
		return fmt.Sprintf("Unknown(%d)", status)
	}
}
