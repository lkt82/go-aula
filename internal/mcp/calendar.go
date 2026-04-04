package aulamcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var listEventsTool = mcp.NewTool("list_events",
	mcp.WithDescription("List calendar events for your children. Defaults to the current week."),
	mcp.WithString("start", mcp.Description("Start date (YYYY-MM-DD)")),
	mcp.WithString("end", mcp.Description("End date (YYYY-MM-DD)")),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var showEventTool = mcp.NewTool("show_event",
	mcp.WithDescription("Show details of a specific calendar event."),
	mcp.WithNumber("event_id", mcp.Description("Event ID"), mcp.Required()),
)

func (s *AulaServer) listEvents(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	children, err := s.childrenIDsByName(ctx, req.GetString("child", ""))
	if err != nil {
		return toolError(err.Error()), nil
	}
	start := req.GetString("start", "")
	end := req.GetString("end", "")
	params := &models.GetEventsParameters{
		InstProfileIDs: children,
	}
	if start != "" {
		params.Start = &start
	}
	if end != "" {
		params.End = &end
	}
	events, err := services.GetEvents(ctx, s.session, params)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list events: %v", err)), nil
	}
	return toolText(formatEvents(events)), nil
}

func (s *AulaServer) showEvent(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	eventID, err := req.RequireInt("event_id")
	if err != nil {
		return toolError("event_id is required"), nil
	}
	detail, err := services.GetEventDetail(ctx, s.session, int64(eventID))
	if err != nil {
		return toolError(fmt.Sprintf("Failed to get event: %v", err)), nil
	}
	return toolText(formatEventDetail(detail)), nil
}

func formatEvents(events []models.EventSimpleDto) string {
	if len(events) == 0 {
		return "No events found."
	}
	var b strings.Builder
	for _, e := range events {
		title := "(untitled)"
		if e.Title != nil {
			title = aulaapi.ExpandTitle(*e.Title)
		}
		id := ""
		if e.ID != nil {
			id = fmt.Sprintf("[%d] ", *e.ID)
		}
		start := ""
		if e.StartDateTime != nil {
			start = *e.StartDateTime
		}
		fmt.Fprintf(&b, "- %s%s (%s)\n", id, title, start)
	}
	return b.String()
}

func formatEventDetail(detail models.EventDetailsDto) string {
	var b strings.Builder
	if detail.Title != nil {
		fmt.Fprintf(&b, "Title: %s\n", aulaapi.ExpandTitle(*detail.Title))
	}
	if detail.StartDateTime != nil {
		fmt.Fprintf(&b, "Start: %s\n", *detail.StartDateTime)
	}
	if detail.EndDateTime != nil {
		fmt.Fprintf(&b, "End: %s\n", *detail.EndDateTime)
	}
	if detail.Description != nil && detail.Description.HTML != nil {
		fmt.Fprintf(&b, "\n%s\n", stripHTML(*detail.Description.HTML))
	}
	return b.String()
}
