package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GetNotifications fetches in-app notifications for the active profile.
func GetNotifications(ctx context.Context, s *aulaapi.Session, childrenIDs []int64, institutionCodes []string) ([]models.NotificationItemDto, error) {
	var query []string
	for _, id := range childrenIDs {
		query = append(query, fmt.Sprintf("activeChildrenIds[]=%d", id))
	}
	for _, code := range institutionCodes {
		query = append(query, fmt.Sprintf("activeInstitutionCodes[]=%s", code))
	}
	path := "?method=notifications.getNotificationsForActiveProfile"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.NotificationItemDto](ctx, s, path)
}

// DeleteNotifications deletes all notifications for the active profile.
func DeleteNotifications(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, "?method=notifications.deleteNotifications")
}

// DeleteNotificationForChild deletes notifications for a specific related child.
func DeleteNotificationForChild(ctx context.Context, s *aulaapi.Session, childInstitutionProfileID int64) (json.RawMessage, error) {
	return aulaapi.SessionGet[json.RawMessage](ctx, s, fmt.Sprintf("?method=notifications.deleteNotificationsByRelatedChild&childInstitutionProfileId=%d", childInstitutionProfileID))
}
