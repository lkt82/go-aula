package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkt82/go-aula/internal/aulaapi"
)

// WidgetTokenResponse is the widget SSO token response.
type WidgetTokenResponse struct {
	Token *string `json:"token,omitempty"`
}

// GetAulaToken gets a widget SSO token for authenticating with embedded widgets.
func GetAulaToken(ctx context.Context, s *aulaapi.Session, widgetID string) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, fmt.Sprintf("?method=aulaToken.getAulaToken&WidgetId=%s", widgetID))
}
