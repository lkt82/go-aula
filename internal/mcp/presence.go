package aulamcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var presenceStatusTool = mcp.NewTool("presence_status",
	mcp.WithDescription("Get current presence status for your children (checked in, checked out, etc)."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var dailyOverviewTool = mcp.NewTool("daily_overview",
	mcp.WithDescription("Get today's presence overview for your children (check-in/out times, location)."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

func (s *AulaServer) presenceStatus(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	children, err := s.childrenIDsByName(ctx, req.GetString("child", ""))
	if err != nil {
		return toolError(err.Error()), nil
	}
	states, err := services.GetChildrensState(ctx, s.session, children)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to get presence: %v", err)), nil
	}
	return toolText(formatPresenceStatus(states)), nil
}

func (s *AulaServer) dailyOverview(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	children, err := s.childrenIDsByName(ctx, req.GetString("child", ""))
	if err != nil {
		return toolError(err.Error()), nil
	}
	regs, err := services.GetDailyOverview(ctx, s.session, children)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to get daily overview: %v", err)), nil
	}
	return toolText(formatDailyOverview(regs)), nil
}

func formatPresenceStatus(states []models.ChildStatusDto) string {
	if len(states) == 0 {
		return "No presence status found."
	}
	var b strings.Builder
	for _, s := range states {
		name := "(unknown)"
		if s.UniStudent != nil && s.UniStudent.Name != nil {
			name = *s.UniStudent.Name
		}
		status := "(unknown)"
		if s.State != nil {
			status = s.State.String()
		}
		fmt.Fprintf(&b, "- %s: %s\n", name, status)
	}
	return b.String()
}

func formatDailyOverview(regs []models.ParentsDailyOverviewResult) string {
	if len(regs) == 0 {
		return "No presence registrations found."
	}
	b, _ := json.MarshalIndent(regs, "", "  ")
	return string(b)
}
