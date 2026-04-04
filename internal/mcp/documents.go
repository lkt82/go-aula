package aulamcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var listDocumentsTool = mcp.NewTool("list_documents",
	mcp.WithDescription("List secure documents."),
	mcp.WithString("child", mcp.Description("Filter by child name (partial match)")),
)

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
