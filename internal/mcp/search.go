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
)

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

func formatSearchResults(result models.SearchResponse) string {
	b, _ := json.MarshalIndent(result, "", "  ")
	return string(b)
}
