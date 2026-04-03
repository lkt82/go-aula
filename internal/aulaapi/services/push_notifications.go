package services

import (
	"context"
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// ClearBadgesRequest is the module identifier for clearing notification badges.
type ClearBadgesRequest struct {
	Module *string `json:"module,omitempty"`
}

// RegisterDevice registers a device for push notifications.
func RegisterDevice(ctx context.Context, s *aulaapi.Session, device *models.ConfigureDeviceModel) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=notifications.registerDevice", device)
}

// UnregisterDevice unregisters a specific device from push notifications.
func UnregisterDevice(ctx context.Context, s *aulaapi.Session, deviceID string) (json.RawMessage, error) {
	body := map[string]string{"deviceId": deviceID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=notifications.unregisterDevice", body)
}

// DeleteAllDevices deletes all registered devices for the current profile.
func DeleteAllDevices(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, "?method=notifications.unregisterAllDevices")
}

// GetNotificationSettings gets the notification settings for the current profile.
func GetNotificationSettings(ctx context.Context, s *aulaapi.Session) (models.NotificationSettings, error) {
	return aulaapi.SessionGet[models.NotificationSettings](ctx, s, "?method=notifications.getNotificationSettingsForActiveProfile&includeDeviceTokens=true")
}

// UpdateNotificationSettings updates notification settings for the current profile.
func UpdateNotificationSettings(ctx context.Context, s *aulaapi.Session, settings *models.NotificationSettings) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=notifications.setNotificationSettingsForActiveProfile", settings)
}

// ClearNotificationBadges clears notification badge counts for a specific module.
func ClearNotificationBadges(ctx context.Context, s *aulaapi.Session, request *ClearBadgesRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=notifications.deleteBadgeNotificationByModule", request)
}
