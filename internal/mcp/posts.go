package aulamcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
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

func (s *AulaServer) listPosts(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if err := s.session.EnsureContextInitialized(ctx); err != nil {
		return toolError(fmt.Sprintf("Failed to initialize session: %v", err)), nil
	}
	profileIDs := s.session.AllInstitutionProfileIDs()
	parent := "profile"
	params := &models.GetPostApiParameters{
		Parent:                &parent,
		InstitutionProfileIDs: profileIDs,
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
