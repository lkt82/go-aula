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

	if err := s.session.EnsureContextInitialized(ctx); err != nil {
		return toolError(fmt.Sprintf("Failed to initialize session: %v", err)), nil
	}

	institutionProfileIDs := s.session.InstitutionProfileIDs()
	childrenProfileIDs := s.session.ChildrenInstProfileIDs()
	limit := 20
	offset := 0

	// Search across common doc types and merge results
	docTypes := []string{"ThreadMessage", "Post", "Profile", "Gallery"}
	var allItems []models.SearchResultItem
	total := 0
	succeeded := 0

	for _, dt := range docTypes {
		docType := dt
		params := &models.GlobalSearchParameters{
			Text:                                &query,
			Limit:                               &limit,
			Offset:                              &offset,
			DocTypeCount:                        true,
			DocType:                             &docType,
			InstitutionProfileIDs:               institutionProfileIDs,
			ActiveChildrenInstitutionProfileIDs: childrenProfileIDs,
		}
		result, err := services.GlobalSearch(ctx, s.session, params)
		if err != nil {
			continue
		}
		succeeded++
		if result.TotalSize != nil {
			total += *result.TotalSize
		}
		allItems = append(allItems, result.Results...)
	}

	if succeeded == 0 {
		return toolError("Search failed: all search endpoints returned errors"), nil
	}

	combined := models.SearchResponse{
		TotalSize: &total,
		Results:   allItems,
	}
	return toolText(formatSearchResults(combined)), nil
}

func formatSearchResults(result models.SearchResponse) string {
	b, _ := json.MarshalIndent(result, "", "  ")
	return string(b)
}
