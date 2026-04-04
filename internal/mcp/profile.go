package aulamcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/mark3labs/mcp-go/mcp"
)

var profileTool = mcp.NewTool("profile",
	mcp.WithDescription("Show your Aula profile information (name, email, phone)."),
)

var listChildrenTool = mcp.NewTool("list_children",
	mcp.WithDescription("List your children with their class/group and institution."),
)

func (s *AulaServer) profile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	prof, err := services.GetProfileMasterData(ctx, s.session)
	if err != nil {
		return toolError(fmt.Sprintf("Failed to get profile: %v", err)), nil
	}
	return toolText(formatProfile(prof)), nil
}

func (s *AulaServer) listChildren(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if err := s.session.EnsureContextInitialized(ctx); err != nil {
		return toolError(fmt.Sprintf("Failed to initialize session: %v", err)), nil
	}
	pd := s.session.ProfileData()
	if pd == nil {
		return toolError("No profile data available"), nil
	}
	var b strings.Builder
	for _, p := range pd.Profiles {
		for _, c := range p.Children {
			name := "(unknown)"
			if c.Name != nil {
				name = *c.Name
			}
			inst := ""
			if c.InstitutionProfile != nil && c.InstitutionProfile.InstitutionName != nil {
				inst = *c.InstitutionProfile.InstitutionName
			}

			// Look up groups to find the class
			var groups []string
			if c.InstitutionProfile != nil && c.InstitutionProfile.ID != nil {
				glist, err := services.GetGroupByContext(ctx, s.session, *c.InstitutionProfile.ID)
				if err == nil {
					for _, g := range glist {
						if g.Name != nil {
							groups = append(groups, *g.Name)
						}
					}
				}
			}

			fmt.Fprintf(&b, "- %s", name)
			if len(groups) > 0 {
				fmt.Fprintf(&b, " (%s)", strings.Join(groups, ", "))
			}
			if inst != "" {
				fmt.Fprintf(&b, " — %s", inst)
			}
			b.WriteString("\n")
		}
	}
	if b.Len() == 0 {
		return toolText("No children found."), nil
	}
	return toolText(b.String()), nil
}

func formatProfile(prof models.Profile) string {
	b, _ := json.MarshalIndent(prof, "", "  ")
	return string(b)
}
