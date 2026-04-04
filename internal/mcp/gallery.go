package aulamcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var listAlbumsTool = mcp.NewTool("list_albums",
	mcp.WithDescription("List photo albums from the gallery."),
)

func (s *AulaServer) listAlbums(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filter := &models.GalleryViewFilter{}
	albums, err := services.GetAlbums(ctx, s.session, filter)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to list albums: %v", err)), nil
	}
	return toolText(formatAlbums(albums)), nil
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
