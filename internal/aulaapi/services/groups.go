package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// JoinOrLeaveGroupRequest is the request body for joining or leaving a group.
type JoinOrLeaveGroupRequest struct {
	Action *string `json:"action,omitempty"`
}

// GetGroup gets a group by its ID.
func GetGroup(ctx context.Context, s *aulaapi.Session, groupID int64) (models.Group, error) {
	return aulaapi.SessionGet[models.Group](ctx, s, fmt.Sprintf("?method=groups.getGroupById&groupId=%d", groupID))
}

// GetGroupByContext gets groups by context (e.g., for a specific institution profile).
func GetGroupByContext(ctx context.Context, s *aulaapi.Session, contextID int64) ([]models.GroupByContextDto, error) {
	return aulaapi.SessionGet[[]models.GroupByContextDto](ctx, s, fmt.Sprintf("?method=groups.getGroupsByContext&contextId=%d", contextID))
}

// GetMembershipsLight gets light membership list for a group.
func GetMembershipsLight(ctx context.Context, s *aulaapi.Session, groupID int64) ([]models.GroupMembership, error) {
	return aulaapi.SessionGet[[]models.GroupMembership](ctx, s, fmt.Sprintf("?method=groups.getMembershipsLight&groupId=%d", groupID))
}

// JoinOrLeaveGroup joins or leaves a group.
func JoinOrLeaveGroup(ctx context.Context, s *aulaapi.Session, groupID int64, request *JoinOrLeaveGroupRequest) (json.RawMessage, error) {
	_ = groupID
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=groups.joinOrLeaveGroup", request)
}
