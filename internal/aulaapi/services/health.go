package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
)

// IsAlive checks if the Aula backend is alive.
func IsAlive(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, "alivecheck/")
}
