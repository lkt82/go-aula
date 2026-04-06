package aulamcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var searchTool = mcp.NewTool("search",
	mcp.WithDescription("Search across Aula content (messages, posts, profiles, groups)."),
	mcp.WithString("query", mcp.Description("Search text"), mcp.Required()),
	mcp.WithString("type", mcp.Description("Content type: ThreadMessage (default), Post, Profile, Event")),
)

func (s *AulaServer) search(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := req.RequireString("query")
	if err != nil {
		return toolError("query is required"), nil
	}

	docType := req.GetString("type", "ThreadMessage")

	if err := s.session.EnsureContextInitialized(ctx); err != nil {
		return toolError(fmt.Sprintf("Failed to initialize session: %v", err)), nil
	}

	limit := 20
	offset := 0
	params := &models.GlobalSearchParameters{
		Text:                                &query,
		Limit:                               &limit,
		Offset:                              &offset,
		DocTypeCount:                        true,
		DocType:                             &docType,
		InstitutionProfileIDs:               s.session.InstitutionProfileIDs(),
		ActiveChildrenInstitutionProfileIDs: s.session.ChildrenInstProfileIDs(),
	}
	result, err := services.GlobalSearch(ctx, s.session, params)
	if err != nil {
		return toolError(fmt.Sprintf("Search failed: %v", err)), nil
	}
	return toolText(formatSearchResults(result)), nil
}

func formatSearchResults(result models.SearchResponse) string {
	b, _ := json.MarshalIndent(result, "", "  ")
	return string(b)
}
