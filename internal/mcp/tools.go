package aulamcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

// ---------------------------------------------------------------------------
// Tool definitions
// ---------------------------------------------------------------------------

var listMessagesTool = mcp.NewTool("list_messages",
	mcp.WithDescription("List message threads from Aula. Returns thread subjects, senders, and read status."),
	mcp.WithNumber("page", mcp.Description("Page number (default 0)")),
)

var readMessageTool = mcp.NewTool("read_message",
	mcp.WithDescription("Read messages in a specific thread by thread ID."),
	mcp.WithNumber("thread_id", mcp.Description("Thread ID to read"), mcp.Required()),
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

var listPostsTool = mcp.NewTool("list_posts",
	mcp.WithDescription("List posts from the institution feed."),
	mcp.WithNumber("page", mcp.Description("Page number (default 0)")),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var showPostTool = mcp.NewTool("show_post",
	mcp.WithDescription("Show a specific post by ID."),
	mcp.WithNumber("post_id", mcp.Description("Post ID"), mcp.Required()),
)

var presenceStatusTool = mcp.NewTool("presence_status",
	mcp.WithDescription("Get current presence status for your children (checked in, checked out, etc)."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var dailyOverviewTool = mcp.NewTool("daily_overview",
	mcp.WithDescription("Get today's presence overview for your children (check-in/out times, location)."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var listNotificationsTool = mcp.NewTool("list_notifications",
	mcp.WithDescription("List recent notifications from Aula."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var searchTool = mcp.NewTool("search",
	mcp.WithDescription("Search across Aula content (messages, posts, profiles, groups)."),
	mcp.WithString("query", mcp.Description("Search text"), mcp.Required()),
)

var listAlbumsTool = mcp.NewTool("list_albums",
	mcp.WithDescription("List photo albums from the gallery."),
)

var listDocumentsTool = mcp.NewTool("list_documents",
	mcp.WithDescription("List secure documents."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

var profileTool = mcp.NewTool("profile",
	mcp.WithDescription("Show your Aula profile information (name, email, phone)."),
)

// ---------------------------------------------------------------------------
// Tool handlers
// ---------------------------------------------------------------------------

func (s *AulaServer) listMessages(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	page := req.GetInt("page", 0)
	args := &models.GetThreadListArguments{
		Page: &page,
	}
	result, err := services.GetThreadList(ctx, s.session, args)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list messages: %v", err)), nil
	}
	return toolText(formatThreadList(result)), nil
}

func (s *AulaServer) readMessage(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	threadID, err := req.RequireInt("thread_id")
	if err != nil {
		return toolError("thread_id is required"), nil
	}
	id := int64(threadID)
	args := &models.GetMessagesForThreadArguments{
		ThreadID: &id,
	}
	result, err := services.GetThreadByID(ctx, s.session, args)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to read thread: %v", err)), nil
	}
	return toolText(formatThread(result)), nil
}

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

func (s *AulaServer) listPosts(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	children, err := s.childrenIDsByName(ctx, req.GetString("child", ""))
	if err != nil {
		return toolError(err.Error()), nil
	}
	params := &models.GetPostApiParameters{
		InstitutionProfileIDs: children,
	}
	result, err := services.GetPosts(ctx, s.session, params)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list posts: %v", err)), nil
	}
	return toolText(formatPosts(result)), nil
}

func (s *AulaServer) showPost(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	postID, err := req.RequireInt("post_id")
	if err != nil {
		return toolError("post_id is required"), nil
	}
	post, err := services.GetPostByID(ctx, s.session, int64(postID))
	if err != nil {
		return toolError(fmt.Sprintf("Failed to get post: %v", err)), nil
	}
	return toolText(formatPostDetail(post)), nil
}

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

func (s *AulaServer) search(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := req.RequireString("query")
	if err != nil {
		return toolError("query is required"), nil
	}
	params := &models.GlobalSearchParameters{
		Text: &query,
	}
	result, err := services.GlobalSearch(ctx, s.session, params)
	if err != nil {
		return toolError(fmt.Sprintf("Search failed: %v", err)), nil
	}
	return toolText(formatSearchResults(result)), nil
}

func (s *AulaServer) listAlbums(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filter := &models.GalleryViewFilter{}
	albums, err := services.GetAlbums(ctx, s.session, filter)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list albums: %v", err)), nil
	}
	return toolText(formatAlbums(albums)), nil
}

func (s *AulaServer) listDocuments(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	children, err := s.childrenIDsByName(ctx, req.GetString("child", ""))
	if err != nil {
		return toolError(err.Error()), nil
	}
	args := &models.GetSecureDocumentsArguments{
		FilterInstitutionProfileIDs: children,
	}
	result, err := services.GetSecureDocuments(ctx, s.session, args)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list documents: %v", err)), nil
	}
	return toolText(formatDocuments(result)), nil
}

func (s *AulaServer) profile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	prof, err := services.GetProfileMasterData(ctx, s.session)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to get profile: %v", err)), nil
	}
	return toolText(formatProfile(prof)), nil
}

// ---------------------------------------------------------------------------
// Formatters — produce LLM-friendly text from API responses
// ---------------------------------------------------------------------------

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
		fmt.Fprintf(&b, "- [%d] %s%s\n", id, subject, unread)
	}
	return b.String()
}

func formatThread(result models.MessagesInThreadDto) string {
	var b strings.Builder
	for _, msg := range result.Messages {
		sender := "Unknown"
		if msg.Sender != nil && msg.Sender.FullName != nil {
			sender = *msg.Sender.FullName
		}
		fmt.Fprintf(&b, "From: %s\n", sender)
		if msg.SendDateTime != nil {
			fmt.Fprintf(&b, "Date: %s\n", *msg.SendDateTime)
		}
		if msg.Text != nil && msg.Text.HTML != nil {
			fmt.Fprintf(&b, "%s\n", stripHTML(*msg.Text.HTML))
		}
		b.WriteString("---\n")
	}
	return b.String()
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

func formatPosts(result models.GetPostApiResult) string {
	if len(result.Posts) == 0 {
		return "No posts found."
	}
	var b strings.Builder
	for _, p := range result.Posts {
		title := "(untitled)"
		if p.Title != nil {
			title = *p.Title
		}
		id := int64(0)
		if p.ID != nil {
			id = *p.ID
		}
		fmt.Fprintf(&b, "- [%d] %s\n", id, title)
	}
	return b.String()
}

func formatPostDetail(post models.PostApiDto) string {
	var b strings.Builder
	if post.Title != nil {
		fmt.Fprintf(&b, "Title: %s\n", *post.Title)
	}
	if post.Content != nil && post.Content.HTML != nil {
		fmt.Fprintf(&b, "\n%s\n", stripHTML(*post.Content.HTML))
	}
	return b.String()
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

func formatSearchResults(result models.SearchResponse) string {
	b, _ := json.MarshalIndent(result, "", "  ")
	return string(b)
}

func formatAlbums(albums []models.AlbumDto) string {
	if len(albums) == 0 {
		return "No albums found."
	}
	var b strings.Builder
	for _, a := range albums {
		title := "(untitled)"
		if a.Title != nil {
			title = *a.Title
		}
		id := int64(0)
		if a.ID != nil {
			id = *a.ID
		}
		fmt.Fprintf(&b, "- [%d] %s\n", id, title)
	}
	return b.String()
}

func formatDocuments(result models.GetSecureDocumentsResult) string {
	if len(result.Documents) == 0 {
		return "No documents found."
	}
	var b strings.Builder
	for _, d := range result.Documents {
		title := "(untitled)"
		if d.Title != nil {
			title = *d.Title
		}
		id := int64(0)
		if d.ID != nil {
			id = *d.ID
		}
		fmt.Fprintf(&b, "- [%d] %s\n", id, title)
	}
	return b.String()
}

func formatProfile(prof models.Profile) string {
	b, _ := json.MarshalIndent(prof, "", "  ")
	return string(b)
}

// stripHTML is a simple HTML tag stripper for LLM-friendly output.
func stripHTML(s string) string {
	var b strings.Builder
	inTag := false
	for _, r := range s {
		if r == '<' {
			inTag = true
			continue
		}
		if r == '>' {
			inTag = false
			continue
		}
		if !inTag {
			b.WriteRune(r)
		}
	}
	return strings.TrimSpace(b.String())
}
