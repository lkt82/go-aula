package aulamcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var listNotificationsTool = mcp.NewTool("list_notifications",
	mcp.WithDescription("List recent notifications from Aula."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

func (s *AulaServer) listNotifications(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	children, err := s.childrenIDsByName(ctx, req.GetString("child", ""))
	if err != nil {
		return toolError(err.Error()), nil
	}
	codes, err := s.institutionCodes(ctx)
	if err != nil {
		return toolError(err.Error()), nil
	}
	notifs, err := services.GetNotifications(ctx, s.session, children, codes)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list notifications: %v", err)), nil
	}
	return toolText(formatNotifications(notifs)), nil
}

func formatNotifications(notifs []models.NotificationItemDto) string {
	if len(notifs) == 0 {
		return "No notifications."
	}
	var b strings.Builder
	for _, n := range notifs {
		title := ""
		if n.Title != nil {
			title = *n.Title
		}
		area := ""
		if n.NotificationArea != nil {
			area = fmt.Sprintf(" [%s]", *n.NotificationArea)
		}
		fmt.Fprintf(&b, "- %s%s\n", title, area)
	}
	return b.String()
}
