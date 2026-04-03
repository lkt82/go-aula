package models

// NotificationItemDto represents an in-app notification.
type NotificationItemDto struct {
	NotificationID                        *string  `json:"notificationId,omitempty"`
	GeneralInformationID                  *int64   `json:"generalInformationId,omitempty"`
	InstitutionProfileID                  *int64   `json:"institutionProfileId,omitempty"`
	NotificationEventType                 *string  `json:"notificationEventType,omitempty"`
	NotificationArea                      *string  `json:"notificationArea,omitempty"`
	NotificationType                      *string  `json:"notificationType,omitempty"`
	InstitutionCode                       *InstitutionCode `json:"institutionCode,omitempty"`
	Expires                               *string  `json:"expires,omitempty"`
	ResponseDeadline                      *string  `json:"responseDeadline,omitempty"`
	Triggered                             *string  `json:"triggered,omitempty"`
	URL                                   *string  `json:"url,omitempty"`
	Content                               *HtmlDto `json:"content,omitempty"`
	RelatedChildInstitutionProfileID      *int64   `json:"relatedChildInstitutionProfileId,omitempty"`
	RelatedChildName                      *string  `json:"relatedChildName,omitempty"`
	Title                                 *string  `json:"title,omitempty"`
	OriginalTitle                         *string  `json:"originalTitle,omitempty"`
	EventID                               *int64   `json:"eventId,omitempty"`
	StartTime                             *string  `json:"startTime,omitempty"`
	EndTime                               *string  `json:"endTime,omitempty"`
	StartDate                             *string  `json:"startDate,omitempty"`
	EndDate                               *string  `json:"endDate,omitempty"`
	OtherCalendarPersonName               *string  `json:"otherCalendarPersonName,omitempty"`
	OtherCalendarInstitutionProfileID     *int     `json:"otherCalendarInstitutionProfileId,omitempty"`
	ResponderName                         *string  `json:"responderName,omitempty"`
	SenderName                            *string  `json:"senderName,omitempty"`
	MessageText                           *string  `json:"messageText,omitempty"`
	RelatedInstitution                    *string  `json:"relatedInstitution,omitempty"`
	FolderID                              *int64   `json:"folderId,omitempty"`
	ThreadID                              *int64   `json:"threadId,omitempty"`
	PostTitle                             *string  `json:"postTitle,omitempty"`
	PostID                                *int64   `json:"postId,omitempty"`
	GroupName                             *string  `json:"groupName,omitempty"`
	GroupID                               *int64   `json:"groupId,omitempty"`
	AlbumID                               *int64   `json:"albumId,omitempty"`
	AlbumName                             *string  `json:"albumName,omitempty"`
	MediaID                               *int64   `json:"mediaId,omitempty"`
	MediaIDs                              []int64  `json:"mediaIds,omitempty"`
	DocumentID                            *int64   `json:"documentId,omitempty"`
	CommonFileID                          *int64   `json:"commonFileId,omitempty"`
	RoomName                              *string  `json:"roomName,omitempty"`
	EventStartTime                        *string  `json:"eventStartTime,omitempty"`
	EventEndTime                          *string  `json:"eventEndTime,omitempty"`
	VacationRegistrationResponseID        *int64   `json:"vacationRegistrationResponseId,omitempty"`
	CommonInboxID                         *int64   `json:"commonInboxId,omitempty"`
	CommonInboxName                       *string  `json:"commonInboxName,omitempty"`
	NoteToGuardians                       *string  `json:"noteToGuardians,omitempty"`
	IsPresenceTimesRequired               bool     `json:"isPresenceTimesRequired"`
	VacationRequestName                   *string  `json:"vacationRequestName,omitempty"`
	NotificationMessage                   *string  `json:"notificationMessage,omitempty"`
	OccurrenceDateTime                    *string  `json:"occurrenceDateTime,omitempty"`
	CancelledBy                           *string  `json:"cancelledBy,omitempty"`
	WidgetID                              *int     `json:"widgetId,omitempty"`
	WidgetName                            *string  `json:"widgetName,omitempty"`
	Message                               *string  `json:"message,omitempty"`
	ResourceName                          *string  `json:"resourceName,omitempty"`
	OccurrenceDate                        *string  `json:"occurrenceDate,omitempty"`
	ExceptionEventID                      *int64   `json:"exceptionEventId,omitempty"`
	CommentID                             *int64   `json:"commentId,omitempty"`
	ProfilePictureInstitutionProfileID    *int64   `json:"profilePictureInstitutionProfileId,omitempty"`
}

// NotificationSettings represents user notification settings.
type NotificationSettings struct {
	ScheduledTime                       *string                       `json:"scheduledTime,omitempty"`
	Instant                             bool                          `json:"instant"`
	Monday                              bool                          `json:"monday"`
	Tuesday                             bool                          `json:"tuesday"`
	Wednesday                           bool                          `json:"wednesday"`
	Thursday                            bool                          `json:"thursday"`
	Friday                              bool                          `json:"friday"`
	Saturday                            bool                          `json:"saturday"`
	Sunday                              bool                          `json:"sunday"`
	NotifyMessages                      bool                          `json:"notifyMessages"`
	NotifyMessagesFromEmployees         bool                          `json:"notifyMessagesFromEmployees"`
	NotifyMessagesFromChildren          bool                          `json:"notifyMessagesFromChildren"`
	NotifyMessagesFromGuardians         bool                          `json:"notifyMessagesFromGuardians"`
	NotifyCalendar                      bool                          `json:"notifyCalendar"`
	NotifyGallery                       bool                          `json:"notifyGallery"`
	NotifyPosts                         bool                          `json:"notifyPosts"`
	EmailDisabled                       bool                          `json:"emailDisabled"`
	EmailAvailable                      bool                          `json:"emailAvailable"`
	AppDisabled                         bool                          `json:"appDisabled"`
	AppAvailable                        bool                          `json:"appAvailable"`
	NotifyAssignedAsSubstituteTeacher   bool                          `json:"notifyAssignedAsSubstituteTeacher"`
	NotifyLessonNoteChanged             bool                          `json:"notifyLessonNoteChanged"`
	ComeGoNotificationSettings          []ComeGoNotificationSettings  `json:"comeGoNotificationSettings,omitempty"`
	DeviceList                          []SimpleDevice                `json:"deviceList,omitempty"`
	WidgetSettings                      []WidgetNotificationSettings  `json:"widgetNotificationSettingDtos,omitempty"`
}

// ComeGoNotificationSettings represents a presence/come-go notification channel setting.
type ComeGoNotificationSettings struct {
	ComeGoType *string `json:"comeGoType,omitempty"`
	Activated  bool    `json:"activated"`
}

// WidgetNotificationSettings represents a per-widget notification setting.
type WidgetNotificationSettings struct {
	Title    *string `json:"title,omitempty"`
	WidgetID *int    `json:"widgetId,omitempty"`
	IsActive bool    `json:"isActive"`
}

// ConfigureDeviceModel represents a device registration payload.
type ConfigureDeviceModel struct {
	CurrentToken        *string `json:"currentToken,omitempty"`
	DeviceID            *string `json:"deviceId,omitempty"`
	DeviceDescription   *string `json:"deviceDescription,omitempty"`
	DeviceAccessGranted bool    `json:"deviceAccessGranted"`
	Platform            *string `json:"platform,omitempty"`
}

// SimpleDevice represents a minimal device identity.
type SimpleDevice struct {
	DeviceID *string `json:"deviceId,omitempty"`
}

// RemoteNotification represents a parsed remote/push notification payload.
type RemoteNotification struct {
	ProfileID                          *int64  `json:"profileId,omitempty"`
	ElementID                          *string `json:"elementId,omitempty"`
	ID                                 *string `json:"id,omitempty"`
	NotificationType                   *string `json:"type,omitempty"`
	RelatedChildInstProfileID          *int64  `json:"relatedChildInstProfileId,omitempty"`
	CommonInboxID                      *int64  `json:"commonInboxId,omitempty"`
	CommentID                          *int    `json:"commentId,omitempty"`
	OccurrenceDateTime                 *string `json:"occurrenceDateTime,omitempty"`
	ProfilePictureInstitutionProfileID *int64  `json:"profilePictureInstitutionProfileId,omitempty"`
}
