package aulamcp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

// currentSchoolYearStart returns August 1 of the current school year.
func currentSchoolYearStart() string {
	now := time.Now()
	year := now.Year()
	if now.Month() < time.August {
		year--
	}
	return fmt.Sprintf("%d-08-01", year)
}

var listMessagesTool = mcp.NewTool("list_messages",
	mcp.WithDescription("List message threads from Aula. Returns thread subjects, senders, and read status. Defaults to current school year."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
	mcp.WithString("since", mcp.Description("Only show messages after this date (YYYY-MM-DD). Defaults to current school year (August 1).")),
)

var readMessageTool = mcp.NewTool("read_message",
	mcp.WithDescription("Read messages in a specific thread by thread ID."),
	mcp.WithNumber("thread_id", mcp.Description("Thread ID to read"), mcp.Required()),
)

func (s *AulaServer) listMessages(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	childName := req.GetString("child", "")
	since := req.GetString("since", currentSchoolYearStart())

	sortOn := "date"
	orderDir := "desc"
	page := 0
	args := &models.GetThreadListArguments{
		Page:           &page,
		SortOn:         &sortOn,
		OrderDirection: &orderDir,
	}
	result, err := services.GetThreadList(ctx, s.session, args)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list messages: %v", err)), nil
	}

	threads := result.Threads
	if childName != "" {
		threads = filterThreadsByChild(threads, childName)
	}
	if since != "" {
		threads = filterThreadsSince(threads, since)
	}
	if len(threads) > 20 {
		threads = threads[:20]
	}
	return toolText(formatThreadList(models.MessageThreadSubscriptionList{Threads: threads})), nil
}

func (s *AulaServer) readMessage(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	threadID, err := req.RequireInt("thread_id")
	if err != nil {
		return toolError("thread_id is required"), nil
	}
	id := int64(threadID)
	page := 0
	args := &models.GetMessagesForThreadArguments{
		ThreadID: &id,
		Page:     &page,
	}
	result, err := services.GetThreadByID(ctx, s.session, args)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to read thread: %v", err)), nil
	}
	return toolText(formatThread(result)), nil
}

func filterThreadsSince(threads []models.MessageThreadSubscription, since string) []models.MessageThreadSubscription {
	var filtered []models.MessageThreadSubscription
	for _, t := range threads {
		date := ""
		if t.LatestMessage != nil && t.LatestMessage.SendDateTime != nil && len(*t.LatestMessage.SendDateTime) >= 10 {
			date = (*t.LatestMessage.SendDateTime)[:10]
		}
		if date >= since {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func filterThreadsByChild(threads []models.MessageThreadSubscription, childName string) []models.MessageThreadSubscription {
	nameLower := strings.ToLower(childName)
	var filtered []models.MessageThreadSubscription
	for _, t := range threads {
		for _, rc := range t.RegardingChildren {
			if rc.DisplayName != nil && strings.Contains(strings.ToLower(*rc.DisplayName), nameLower) {
				filtered = append(filtered, t)
				break
			}
		}
	}
	return filtered
}

func formatThreadList(result models.MessageThreadSubscriptionList) string {
	if len(result.Threads) == 0 {
		return "No message threads found."
	}
	var b strings.Builder
	for _, t := range result.Threads {
		subject := "(no subject)"
		if t.Subject != nil {
			subject = *t.Subject
		}
		id := int64(0)
		if t.ID != nil {
			id = *t.ID
		}
		unread := ""
		if !t.Read {
			unread = " [UNREAD]"
		}
		sender := ""
		if t.Creator != nil && t.Creator.FullName != nil {
			sender = fmt.Sprintf(" from %s", *t.Creator.FullName)
		}
		date := ""
		if t.LatestMessage != nil && t.LatestMessage.SendDateTime != nil {
			date = fmt.Sprintf(" (%s)", (*t.LatestMessage.SendDateTime)[:10])
		}
		fmt.Fprintf(&b, "- [%d] %s%s%s%s\n", id, subject, sender, date, unread)
	}
	return b.String()
}

func formatThread(result models.MessagesInThreadDto) string {
	var b strings.Builder
	if result.Subject != nil {
		fmt.Fprintf(&b, "Subject: %s\n\n", *result.Subject)
	}
	for _, msg := range result.Messages {
		if msg.Text == nil || msg.Text.HTML == nil {
			continue
		}
		sender := ""
		if msg.Sender != nil {
			if msg.Sender.FullName != nil {
				sender = *msg.Sender.FullName
			} else if msg.Sender.ShortName != nil {
				sender = *msg.Sender.ShortName
			}
		}
		date := ""
		if msg.SendDateTime != nil {
			date = *msg.SendDateTime
		}
		if sender != "" {
			fmt.Fprintf(&b, "From: %s (%s)\n", sender, date)
		} else {
			fmt.Fprintf(&b, "Date: %s\n", date)
		}
		fmt.Fprintf(&b, "%s\n\n---\n\n", stripHTML(*msg.Text.HTML))
	}
	return b.String()
}
