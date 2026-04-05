package aulamcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/mark3labs/mcp-go/server"
)

// AulaServer wraps an Aula session and exposes it as MCP tools.
type AulaServer struct {
	session *aulaapi.Session
}

// NewAulaServer creates a new MCP server backed by the given Aula session.
func NewAulaServer(session *aulaapi.Session) *server.MCPServer {
	s := &AulaServer{session: session}

	mcpServer := server.NewMCPServer("aula-mcp", "0.1.0")

	mcpServer.AddTools(
		// Messages
		server.ServerTool{Tool: listMessagesTool, Handler: s.listMessages},
		server.ServerTool{Tool: readMessageTool, Handler: s.readMessage},

		// Calendar
		server.ServerTool{Tool: listEventsTool, Handler: s.listEvents},
		server.ServerTool{Tool: showEventTool, Handler: s.showEvent},

		// Posts
		server.ServerTool{Tool: listPostsTool, Handler: s.listPosts},
		server.ServerTool{Tool: showPostTool, Handler: s.showPost},

		// Presence
		server.ServerTool{Tool: presenceStatusTool, Handler: s.presenceStatus},
		server.ServerTool{Tool: dailyOverviewTool, Handler: s.dailyOverview},

		// Notifications
		server.ServerTool{Tool: listNotificationsTool, Handler: s.listNotifications},

		// Search
		server.ServerTool{Tool: searchTool, Handler: s.search},

		// Gallery
		server.ServerTool{Tool: listAlbumsTool, Handler: s.listAlbums},

		// Documents
		server.ServerTool{Tool: listDocumentsTool, Handler: s.listDocuments},

		// Profile
		server.ServerTool{Tool: profileTool, Handler: s.profile},
		server.ServerTool{Tool: listChildrenTool, Handler: s.listChildren},
	)

	return mcpServer
}

// childrenIDsByName resolves children IDs, optionally filtered by name substring.
func (s *AulaServer) childrenIDsByName(ctx context.Context, childName string) ([]int64, error) {
	if err := s.session.EnsureContextInitialized(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize session: %w", err)
	}
	if childName == "" {
		ids := s.session.ChildrenInstProfileIDs()
		if len(ids) == 0 {
			return nil, fmt.Errorf("no children found in profile")
		}
		return ids, nil
	}
	pd := s.session.ProfileData()
	if pd == nil {
		return nil, fmt.Errorf("no profile data available")
	}
	nameLower := strings.ToLower(childName)
	var ids []int64
	for _, p := range pd.Profiles {
		for _, c := range p.Children {
			if c.Name != nil && strings.Contains(strings.ToLower(*c.Name), nameLower) {
				if c.InstitutionProfile != nil && c.InstitutionProfile.ID != nil {
					ids = append(ids, *c.InstitutionProfile.ID)
				}
			}
		}
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("no child found matching '%s'", childName)
	}
	return ids, nil
}

// institutionCodes resolves the institution codes from the session.
func (s *AulaServer) institutionCodes(ctx context.Context) ([]string, error) {
	if err := s.session.EnsureContextInitialized(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize session: %w", err)
	}
	codes := s.session.ChildrenInstitutionCodes()
	return codes, nil
}


