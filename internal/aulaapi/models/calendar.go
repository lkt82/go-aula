package models

import "encoding/json"

// HtmlDto represents a simple HTML content wrapper.
type HtmlDto struct {
	HTML *string `json:"html,omitempty"`
}

// CalendarAulaFileResultDto represents a file attachment result DTO for calendar.
type CalendarAulaFileResultDto struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// EventResourceCategory represents a resource category for event resources.
type EventResourceCategory struct {
	ResourceType *string `json:"resourceType,omitempty"`
}

// EventResource represents a resource attached to a calendar event.
type EventResource struct {
	ID              *int                    `json:"id,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Label           *string                 `json:"label,omitempty"`
	Value           *string                 `json:"value,omitempty"`
	Category        *EventResourceCategory  `json:"category,omitempty"`
	InstitutionCode *string                 `json:"institutionCode,omitempty"`
	InstitutionName *string                 `json:"institutionName,omitempty"`
	ShortName       *string                 `json:"shortName,omitempty"`
}

// EventGroupWithRolesDto represents a group with invited portal roles on an event.
type EventGroupWithRolesDto struct {
	Group              *EventGroup `json:"group,omitempty"`
	InvitedPortalRoles []string    `json:"invitedPortalRoles,omitempty"`
}

// RepeatingEventDto represents repeating event pattern information.
type RepeatingEventDto struct {
	Pattern               *string `json:"pattern,omitempty"`
	OccurenceLimit        *int    `json:"occurenceLimit,omitempty"`
	Interval              *int    `json:"interval,omitempty"`
	DayInMonth            *int    `json:"dayInMonth,omitempty"`
	MaxDate               *string `json:"maxDate,omitempty"`
	WeekdayMask           []bool  `json:"weekdayMask,omitempty"`
	OriginalStartDateTime *string `json:"originalStartDateTime,omitempty"`
	OriginalEndDateTime   *string `json:"originalEndDateTime,omitempty"`
	LastOccurrenceDate    *string `json:"lastOccurrenceDate,omitempty"`
}

// EventBaseClass represents a base class for all calendar events.
type EventBaseClass struct {
	ID                          *int                      `json:"id,omitempty"`
	Title                       *string                   `json:"title,omitempty"`
	EventType                   *string                   `json:"type,omitempty"`
	InstitutionCode             *string                   `json:"institutionCode,omitempty"`
	AllDay                      *bool                     `json:"allDay,omitempty"`
	OldAllDay                   *bool                     `json:"oldAllDay,omitempty"`
	AddedToInstitutionCalendar  bool                      `json:"addedToInstitutionCalendar"`
	HideInOwnCalendar           bool                      `json:"hideInOwnCalendar"`
	ResponseDeadline            *string                   `json:"responseDeadline,omitempty"`
	IsDeadlineExceeded          bool                      `json:"isDeadlineExceeded"`
	StartDateTime               *string                   `json:"startDateTime,omitempty"`
	EndDateTime                 *string                   `json:"endDateTime,omitempty"`
	Private                     *bool                     `json:"private,omitempty"`
	ResponseRequired            *bool                     `json:"responseRequired,omitempty"`
	BelongsToProfiles           []int64                   `json:"belongsToProfiles,omitempty"`
	BelongsToResources          []int64                   `json:"belongsToResources,omitempty"`
	SecurityLevel               *int                      `json:"securityLevel,omitempty"`
	IsDeleted                   bool                      `json:"isDeleted"`
	OldStartDateTime            *string                   `json:"oldStartDateTime,omitempty"`
	OldEndDateTime              *string                   `json:"oldEndDateTime,omitempty"`
	InviteeGroups               []EventGroupWithRolesDto  `json:"inviteeGroups,omitempty"`
	InvitedGroups               []SimpleGroupDto          `json:"invitedGroups,omitempty"`
	PrimaryResourceText         *string                   `json:"primaryResourceText,omitempty"`
	PrimaryResource             *EventResource            `json:"primaryResource,omitempty"`
	AdditionalResources         []EventResource           `json:"additionalResources,omitempty"`
	AdditionalResourceText      *string                   `json:"additionalResourceText,omitempty"`
	Repeating                   *RepeatingEventDto        `json:"repeating,omitempty"`
	ResponseStatus              *string                   `json:"responseStatus,omitempty"`
	DirectlyRelated             bool                      `json:"directlyRelated"`
	MaximumNumberOfParticipants *int64                    `json:"maximumNumberOfParticipants,omitempty"`
	ActualNumberOfParticipants  *int64                    `json:"actualNumberOfParticipants,omitempty"`
	OccurrenceDateTime          *string                   `json:"occurrenceDateTime,omitempty"`
}

// EventDetailsCreatorDto represents creator information on an event detail view.
type EventDetailsCreatorDto struct {
	InstitutionProfileID *int64  `json:"institutionProfileId,omitempty"`
	ProfileID            *int64  `json:"profileId,omitempty"`
	Name                 *string `json:"name,omitempty"`
	ShortName            *string `json:"shortName,omitempty"`
	Metadata             *string `json:"metadata,omitempty"`
	Role                 *string `json:"role,omitempty"`
	InstitutionCode      *string `json:"institutionCode,omitempty"`
	InstitutionName      *string `json:"institutionName,omitempty"`
	InstitutionRole      *string `json:"institutionRole,omitempty"`
}

// EventDetailsDto represents a detailed event view.
type EventDetailsDto struct {
	ID                            *int                                    `json:"id,omitempty"`
	Title                         *string                                 `json:"title,omitempty"`
	EventType                     *string                                 `json:"type,omitempty"`
	InstitutionCode               *string                                 `json:"institutionCode,omitempty"`
	AllDay                        *bool                                   `json:"allDay,omitempty"`
	OldAllDay                     *bool                                   `json:"oldAllDay,omitempty"`
	AddedToInstitutionCalendar    bool                                    `json:"addedToInstitutionCalendar"`
	HideInOwnCalendar             bool                                    `json:"hideInOwnCalendar"`
	ResponseDeadline              *string                                 `json:"responseDeadline,omitempty"`
	IsDeadlineExceeded            bool                                    `json:"isDeadlineExceeded"`
	StartDateTime                 *string                                 `json:"startDateTime,omitempty"`
	EndDateTime                   *string                                 `json:"endDateTime,omitempty"`
	Private                       *bool                                   `json:"private,omitempty"`
	ResponseRequired              *bool                                   `json:"responseRequired,omitempty"`
	BelongsToProfiles             []int64                                 `json:"belongsToProfiles,omitempty"`
	BelongsToResources            []int64                                 `json:"belongsToResources,omitempty"`
	SecurityLevel                 *int                                    `json:"securityLevel,omitempty"`
	IsDeleted                     bool                                    `json:"isDeleted"`
	OldStartDateTime              *string                                 `json:"oldStartDateTime,omitempty"`
	OldEndDateTime                *string                                 `json:"oldEndDateTime,omitempty"`
	InviteeGroups                 []EventGroupWithRolesDto                `json:"inviteeGroups,omitempty"`
	InvitedGroups                 []SimpleGroupDto                        `json:"invitedGroups,omitempty"`
	PrimaryResourceText           *string                                 `json:"primaryResourceText,omitempty"`
	PrimaryResource               *EventResource                          `json:"primaryResource,omitempty"`
	AdditionalResources           []EventResource                         `json:"additionalResources,omitempty"`
	AdditionalResourceText        *string                                 `json:"additionalResourceText,omitempty"`
	Repeating                     *RepeatingEventDto                      `json:"repeating,omitempty"`
	ResponseStatus                *string                                 `json:"responseStatus,omitempty"`
	DirectlyRelated               bool                                    `json:"directlyRelated"`
	MaximumNumberOfParticipants   *int64                                  `json:"maximumNumberOfParticipants,omitempty"`
	ActualNumberOfParticipants    *int64                                  `json:"actualNumberOfParticipants,omitempty"`
	OccurrenceDateTime            *string                                 `json:"occurrenceDateTime,omitempty"`
	Attachments                   []FilesAulaFileResultDto                `json:"attachments,omitempty"`
	Invitees                      []EventProfile                          `json:"invitees,omitempty"`
	CoOrganizers                  []EventProfile                          `json:"coOrganizers,omitempty"`
	InvitedGroupHomeChildren      []InvitedGroupHome                      `json:"invitedGroupHomeChildren,omitempty"`
	Description                   *RichTextWrapperDto                     `json:"description,omitempty"`
	Lesson                        *Lesson                                 `json:"lesson,omitempty"`
	Creator                       *EventDetailsCreatorDto                 `json:"creator,omitempty"`
	VacationRegistration          *VacationRegistrationDetailsResultDto   `json:"vacationRegistration,omitempty"`
	TimeSlot                      *TimeslotEventDto                       `json:"timeSlot,omitempty"`
	CanEditStartDate              *bool                                   `json:"canEditStartDate,omitempty"`
	CanAnswerForSeries            *bool                                   `json:"canAnswerForSeries,omitempty"`
	DoRequestNumberOfParticipants bool                                    `json:"doRequestNumberOfParticipants"`
	LastReminderDateTime          *string                                 `json:"lastReminderDateTime,omitempty"`
}

// EventSimpleDto represents a summary event DTO (list views).
type EventSimpleDto struct {
	ID                            *int                                          `json:"id,omitempty"`
	Title                         *string                                       `json:"title,omitempty"`
	EventType                     *string                                       `json:"type,omitempty"`
	InstitutionCode               *string                                       `json:"institutionCode,omitempty"`
	AllDay                        *bool                                         `json:"allDay,omitempty"`
	AddedToInstitutionCalendar    bool                                          `json:"addedToInstitutionCalendar"`
	HideInOwnCalendar             bool                                          `json:"hideInOwnCalendar"`
	ResponseDeadline              *string                                       `json:"responseDeadline,omitempty"`
	IsDeadlineExceeded            bool                                          `json:"isDeadlineExceeded"`
	StartDateTime                 *string                                       `json:"startDateTime,omitempty"`
	EndDateTime                   *string                                       `json:"endDateTime,omitempty"`
	Private                       *bool                                         `json:"private,omitempty"`
	ResponseRequired              *bool                                         `json:"responseRequired,omitempty"`
	BelongsToProfiles             []int64                                       `json:"belongsToProfiles,omitempty"`
	BelongsToResources            []int64                                       `json:"belongsToResources,omitempty"`
	SecurityLevel                 *int                                          `json:"securityLevel,omitempty"`
	IsDeleted                     bool                                          `json:"isDeleted"`
	OldStartDateTime              *string                                       `json:"oldStartDateTime,omitempty"`
	OldEndDateTime                *string                                       `json:"oldEndDateTime,omitempty"`
	InviteeGroups                 []EventGroupWithRolesDto                      `json:"inviteeGroups,omitempty"`
	InvitedGroups                 []SimpleGroupDto                              `json:"invitedGroups,omitempty"`
	PrimaryResourceText           *string                                       `json:"primaryResourceText,omitempty"`
	PrimaryResource               *EventResource                                `json:"primaryResource,omitempty"`
	AdditionalResources           []EventResource                               `json:"additionalResources,omitempty"`
	AdditionalResourceText        *string                                       `json:"additionalResourceText,omitempty"`
	Repeating                     *RepeatingEventDto                            `json:"repeating,omitempty"`
	ResponseStatus                *string                                       `json:"responseStatus,omitempty"`
	DirectlyRelated               bool                                          `json:"directlyRelated"`
	MaximumNumberOfParticipants   *int64                                        `json:"maximumNumberOfParticipants,omitempty"`
	ActualNumberOfParticipants    *int64                                        `json:"actualNumberOfParticipants,omitempty"`
	OccurrenceDateTime            *string                                       `json:"occurrenceDateTime,omitempty"`
	HasAttachments                *bool                                         `json:"hasAttachments,omitempty"`
	Lesson                        *LessonSimple                                 `json:"lesson,omitempty"`
	VacationChildrenCountByDates  []VacationRegistrationChildrenCountByDates    `json:"vacationChildrenCountByDates,omitempty"`
	CreatorAulaName               *string                                       `json:"creatorAulaName,omitempty"`
	CreatorProfileID              *int64                                        `json:"creatorProfileId,omitempty"`
	CreatorInstProfileID          *int64                                        `json:"creatorInstProfileId,omitempty"`
	TimeSlot                      *TimeslotEventSimpleDto                       `json:"timeSlot,omitempty"`
}

// EventProfileDetails represents profile details within an event context.
type EventProfileDetails struct {
	Email                                       *string `json:"email,omitempty"`
	Administrator                               json.RawMessage `json:"administrator,omitempty"`
	FirstName                                   *string `json:"firstName,omitempty"`
	LastName                                    *string `json:"lastName,omitempty"`
	FullName                                    *string `json:"fullName,omitempty"`
	ShortName                                   *string `json:"shortName,omitempty"`
	Metadata                                    *string `json:"metadata,omitempty"`
	Role                                        *string `json:"role,omitempty"`
	Phone                                       *string `json:"phone,omitempty"`
	CanRemoveBlockingOrResponseForTimeSlot      bool    `json:"canRemoveBlockingOrResponseForTimeSlot"`
	ProfileID                                   *int    `json:"profileId,omitempty"`
	InstitutionProfileID                        *int64  `json:"institutionProfileId,omitempty"`
	ProfilePictureURL                           *string `json:"profilePictureUrl,omitempty"`
}

// EventProfile represents a participant profile on an event.
type EventProfile struct {
	InstProfile                *EventProfileDetails `json:"instProfile,omitempty"`
	ResponseType               *string              `json:"responseType,omitempty"`
	ResponseDateTime           *string              `json:"responseDateTime,omitempty"`
	NumberOfAdultParticipants  *int                 `json:"numberOfAdultParticipants,omitempty"`
	NumberOfChildParticipants  *int                 `json:"numberOfChildParticipants,omitempty"`
}

// EventGroup represents a group associated with an event.
type EventGroup struct {
	ID               *int64  `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	GroupType        *string `json:"type,omitempty"`
	Access           *string `json:"access,omitempty"`
	Status           *string `json:"status,omitempty"`
	DashboardEnabled bool    `json:"dashboardEnabled"`
	InstitutionCode  *string `json:"institutionCode,omitempty"`
}

// InvitedGroupHome represents a group-home child invited to an event.
type InvitedGroupHome struct {
	OtpInboxID                    *int64  `json:"otpInboxId,omitempty"`
	GroupHomeID                   *int64  `json:"groupHomeId,omitempty"`
	RegardingChildID              *int64  `json:"regardingChildId,omitempty"`
	RegardingChildDisplayName     *string `json:"regardingChildDisplayName,omitempty"`
	RegardingChildMetaData        *string `json:"regardingChildMetaData,omitempty"`
	GroupHomeName                 *string `json:"groupHomeName,omitempty"`
	ResponseType                  *string `json:"responseType,omitempty"`
	ResponseDateTime              *string `json:"responseDateTime,omitempty"`
	NumberOfAdultParticipants     *int    `json:"numberOfAdultParticipants,omitempty"`
	NumberOfChildParticipants     *int    `json:"numberOfChildParticipants,omitempty"`
}

// BaseTimeslotEventDto represents the base timeslot event DTO.
type BaseTimeslotEventDto struct {
	ChildRequired *bool `json:"childRequired,omitempty"`
}

// TimeslotEventDto represents a full timeslot event (detail view).
type TimeslotEventDto struct {
	ChildRequired                    *bool      `json:"childRequired,omitempty"`
	MeetingsBetweenBreaks            *int       `json:"meetingsBetweenBreaks,omitempty"`
	BreakLength                      *int       `json:"breakLength,omitempty"`
	MeetingDuration                  *int       `json:"meetingDuration,omitempty"`
	CanUpdateResponseToEvent         bool       `json:"canUpdateResponseToEvent"`
	TimeSlots                        []TimeSlot `json:"timeSlots,omitempty"`
	NumberOfParticipantsPerTimeSlot  *int       `json:"numberOfParticipantsPerTimeSlot,omitempty"`
}

// TimeslotEventSimpleDto represents a simplified timeslot event (list view).
type TimeslotEventSimpleDto struct {
	ChildRequired *bool              `json:"childRequired,omitempty"`
	TimeSlots     []TimeSlotSimpleDto `json:"timeSlots,omitempty"`
}

// TimeSlotIndex represents a time index within a timeslot.
type TimeSlotIndex struct {
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
}

// BaseTimeSlotDto represents the base time slot DTO.
type BaseTimeSlotDto struct {
	ID              *int64          `json:"id,omitempty"`
	StartDate       *string         `json:"startDate,omitempty"`
	EndDate         *string         `json:"endDate,omitempty"`
	TimeSlotIndexes []TimeSlotIndex `json:"timeSlotIndexes,omitempty"`
}

// BaseTimeSlotAnswer represents a base answer to a timeslot booking.
type BaseTimeSlotAnswer struct {
	ID                    *int `json:"id,omitempty"`
	SelectedTimeSlotIndex *int `json:"selectedTimeSlotIndex,omitempty"`
}

// TimeSlotAnswer represents a full timeslot answer (with profile details).
type TimeSlotAnswer struct {
	ID                                        *int                  `json:"id,omitempty"`
	SelectedTimeSlotIndex                     *int                  `json:"selectedTimeSlotIndex,omitempty"`
	ConcerningProfile                         *EventProfileDetails  `json:"concerningProfile,omitempty"`
	InstProfile                               *EventProfileDetails  `json:"instProfile,omitempty"`
	CanRemoveBlockingOrResponseForTimeSlot    bool                  `json:"canRemoveBlockingOrResponseForTimeSlot"`
}

// TimeSlotAnswerSimpleDto represents a simplified timeslot answer.
type TimeSlotAnswerSimpleDto struct {
	ID                                     *int `json:"id,omitempty"`
	SelectedTimeSlotIndex                  *int `json:"selectedTimeSlotIndex,omitempty"`
	ConcerningProfileID                    *int `json:"concerningProfileId,omitempty"`
	InstProfileID                          *int `json:"instProfileId,omitempty"`
	CanRemoveBlockingOrResponseForTimeSlot bool `json:"canRemoveBlockingOrResponseForTimeSlot"`
}

// TimeSlot represents a full timeslot (detail view).
type TimeSlot struct {
	ID                  *int64           `json:"id,omitempty"`
	StartDate           *string          `json:"startDate,omitempty"`
	EndDate             *string          `json:"endDate,omitempty"`
	TimeSlotIndexes     []TimeSlotIndex  `json:"timeSlotIndexes,omitempty"`
	Answers             []TimeSlotAnswer `json:"answers,omitempty"`
	PrimaryResource     *EventResource   `json:"primaryResource,omitempty"`
	PrimaryResourceText *string          `json:"primaryResourceText,omitempty"`
}

// TimeSlotSimpleDto represents a simplified timeslot (list view).
type TimeSlotSimpleDto struct {
	ID                *int64                    `json:"id,omitempty"`
	StartDate         *string                   `json:"startDate,omitempty"`
	EndDate           *string                   `json:"endDate,omitempty"`
	TimeSlotIndexes   []TimeSlotIndex           `json:"timeSlotIndexes,omitempty"`
	Answers           []TimeSlotAnswerSimpleDto `json:"answers,omitempty"`
	BelongsToResource *int64                    `json:"belongsToResource,omitempty"`
}

// LessonBase represents base lesson fields.
type LessonBase struct {
	LessonID     *string `json:"lessonId,omitempty"`
	LessonStatus *string `json:"lessonStatus,omitempty"`
}

// LessonParticipant represents a lesson participant with role and teacher info.
type LessonParticipant struct {
	ParticipantProfile *InstitutionProfile `json:"participantProfile,omitempty"`
	ParticipantRole    *string             `json:"participantRole,omitempty"`
	TeacherName        *string             `json:"teacherName,omitempty"`
	TeacherInitials    *string             `json:"teacherInitials,omitempty"`
}

// Lesson represents a full lesson detail.
type Lesson struct {
	LessonID         *string             `json:"lessonId,omitempty"`
	LessonStatus     *string             `json:"lessonStatus,omitempty"`
	Participants     []LessonParticipant `json:"participants,omitempty"`
	NoteToClass      *HtmlDto            `json:"noteToClass,omitempty"`
	NoteToSubstitute *HtmlDto            `json:"noteToSubstitute,omitempty"`
	NoteToTeacher    *HtmlDto            `json:"noteToTeacher,omitempty"`
}

// ParticipantSimple represents simplified participant info (list views).
type ParticipantSimple struct {
	TeacherInitials *string `json:"teacherInitials,omitempty"`
	TeacherName     *string `json:"teacherName,omitempty"`
	ParticipantRole *string `json:"participantRole,omitempty"`
}

// LessonSimple represents a simplified lesson (list views).
type LessonSimple struct {
	LessonID        *string             `json:"lessonId,omitempty"`
	LessonStatus    *string             `json:"lessonStatus,omitempty"`
	HasRelevantNote bool                `json:"hasRelevantNote"`
	Participants    []ParticipantSimple `json:"participants,omitempty"`
}

// DelegateAccessInstitution represents institution info for delegate access context.
type DelegateAccessInstitution struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
}

// DelegateAccessesItem represents a single delegate access item.
type DelegateAccessesItem struct {
	InstProfileID   *int64  `json:"instProfileId,omitempty"`
	ProfileID       *int64  `json:"profileId,omitempty"`
	Name            *string `json:"name,omitempty"`
	InstitutionCode *string `json:"institutionCode,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
	MetaData        *string `json:"metaData,omitempty"`
}

// DelegateAccesses represents delegated calendar access configuration.
type DelegateAccesses struct {
	OwnerInstProfileID   *int64                 `json:"ownerInstProfileId,omitempty"`
	DelegatedInstProfiles []DelegateAccessesItem `json:"delegatedInstProfiles,omitempty"`
}

// DelegateAccessesInput represents input for setting delegated access.
type DelegateAccessesInput struct {
	OwnerInstProfileID     *int64  `json:"ownerInstProfileId,omitempty"`
	DelegatedInstProfileIDs []int64 `json:"delegatedInstProfileIds,omitempty"`
}

// InstitutionDelegateAccessesItem represents a delegate access item with full institution info.
type InstitutionDelegateAccessesItem struct {
	InstProfileID *int64                     `json:"instProfileId,omitempty"`
	ProfileID     *int64                     `json:"profileId,omitempty"`
	FirstName     *string                    `json:"firstName,omitempty"`
	LastName      *string                    `json:"lastName,omitempty"`
	MetaData      *string                    `json:"metaData,omitempty"`
	Institution   *DelegateAccessInstitution `json:"institution,omitempty"`
}

// DelegatedContextResultModel represents a delegated context result.
type DelegatedContextResultModel struct {
	FullName                        *string            `json:"fullName,omitempty"`
	NullableInstitutionProfileID    *int               `json:"nullableInstitutionProfileId,omitempty"`
	InstitutionCode                 *string            `json:"institutionCode,omitempty"`
	ProfilePicture                  *ProfilePictureDto `json:"profilePicture,omitempty"`
	Role                            *string            `json:"role,omitempty"`
}

// SetDelegatedContextRequestModel represents a set delegated context request.
type SetDelegatedContextRequestModel struct {
	DelegatedInstProfileID *int64 `json:"delegatedInstProfileId,omitempty"`
}

// CalendarSynchronisationConfigurationItem represents a calendar feed sync configuration item.
type CalendarSynchronisationConfigurationItem struct {
	InstitutionProfileID          *int     `json:"institutionProfileId,omitempty"`
	ID                            *int     `json:"id,omitempty"`
	Calendarfeedconfigurationid   *int     `json:"calendarfeedconfigurationid,omitempty"`
	OwnerID                       *int     `json:"ownerId,omitempty"`
	RegardingID                   *int     `json:"regardingId,omitempty"`
	OneWeekFeed                   *string  `json:"oneWeekFeed,omitempty"`
	OneYearFeed                   *string  `json:"oneYearFeed,omitempty"`
	Weekly                        bool     `json:"weekly"`
	Filters                       []string `json:"filters,omitempty"`
	FeedStatus                    *string  `json:"feedStatus,omitempty"`
}

// CalendarSynchronisationModel represents calendar sync policy acceptance.
type CalendarSynchronisationModel struct {
	PolicyAccepted bool `json:"policyAccepted"`
}

// CalendarSynchronisationMunicipalityFeedModel represents municipality-level calendar feed configuration.
type CalendarSynchronisationMunicipalityFeedModel struct {
	MunicipalityCode    *string `json:"municipalityCode,omitempty"`
	CalendarFeedEnabled bool    `json:"calendarFeedEnabled"`
}

// CreateCalendarSynchronizationConfigurationRequest represents creating a new calendar sync configuration.
type CreateCalendarSynchronizationConfigurationRequest struct {
	Filters              []string `json:"filters,omitempty"`
	Weekly               bool     `json:"weekly"`
	InstitutionProfileID *int64   `json:"institutionProfileId,omitempty"`
}

// UpdateCalendarSynchronizationConfigurationRequest represents updating an existing calendar sync configuration.
type UpdateCalendarSynchronizationConfigurationRequest struct {
	Filters                      []string `json:"filters,omitempty"`
	CalendarFeedConfigurationID  *int64   `json:"calendarFeedConfigurationId,omitempty"`
}

// GetEventTypesByPortalRoleResultModel represents event types available by portal role.
type GetEventTypesByPortalRoleResultModel struct {
	EventTypes []string `json:"eventTypes,omitempty"`
}

// BirthdayEventDto represents a birthday event DTO.
type BirthdayEventDto struct {
	Birthday             *string `json:"birthday,omitempty"`
	Name                 *string `json:"name,omitempty"`
	InstitutionCode      *string `json:"institutionCode,omitempty"`
	InstitutionProfileID *int64  `json:"institutionProfileId,omitempty"`
	MainGroupName        *string `json:"mainGroupName,omitempty"`
}

// GetBirthdayEvents represents a request to get birthday events.
type GetBirthdayEvents struct {
	Start     *string  `json:"start,omitempty"`
	End       *string  `json:"end,omitempty"`
	InstCodes []string `json:"instCodes,omitempty"`
}

// ImportantDateItemProfile represents a profile on an important date invitee.
type ImportantDateItemProfile struct {
	ID        *int64                     `json:"id,omitempty"`
	ProfileID *int                       `json:"profileId,omitempty"`
	Role      *string                    `json:"role,omitempty"`
	Relations []ImportantDateItemProfile `json:"relations,omitempty"`
}

// ImportantDateItemInvitee represents an invitee on an important date.
type ImportantDateItemInvitee struct {
	InstProfile  *ImportantDateItemProfile `json:"instProfile,omitempty"`
	ResponseType *string                   `json:"responseType,omitempty"`
}

// ImportantDateItem represents an important date item.
type ImportantDateItem struct {
	ID              *int64                     `json:"id,omitempty"`
	StartDateTime   *string                    `json:"startDateTime,omitempty"`
	EndDateTime     *string                    `json:"endDateTime,omitempty"`
	Title           *string                    `json:"title,omitempty"`
	ItemType        *string                    `json:"type,omitempty"`
	Invitees        []ImportantDateItemInvitee `json:"invitees,omitempty"`
	InstitutionName *string                    `json:"institutionName,omitempty"`
	AllDay          bool                       `json:"allDay"`
}

// MyCalendarItem represents a my calendar item.
type MyCalendarItem struct {
	ItemType            *string         `json:"type,omitempty"`
	MyCalendarViewModel json.RawMessage `json:"myCalendarViewModel,omitempty"`
	Title               *string         `json:"title,omitempty"`
	ID                  *int64          `json:"id,omitempty"`
	BirthDay            *string         `json:"birthDay,omitempty"`
	Name                *string         `json:"name,omitempty"`
	GroupName           *string         `json:"groupName,omitempty"`
}

// EventItem represents a calendar event item.
type EventItem struct {
	ItemType       *string         `json:"type,omitempty"`
	EventViewModel json.RawMessage `json:"eventViewModel,omitempty"`
	Title          *string         `json:"title,omitempty"`
	ID             *int64          `json:"id,omitempty"`
	DateTime       *string         `json:"dateTime,omitempty"`
}

// AggregatedEventsGroupByType represents an aggregated event count by type.
type AggregatedEventsGroupByType struct {
	EventType *string `json:"type,omitempty"`
	Count     *int    `json:"count,omitempty"`
}

// DailyAggregatedEventsResultModel represents a daily aggregated events result.
type DailyAggregatedEventsResultModel struct {
	Date             *string                       `json:"date,omitempty"`
	AggregatedEvents []AggregatedEventsGroupByType `json:"aggregatedEvents,omitempty"`
}

// DailyEventCountResultModel represents a daily event count result.
type DailyEventCountResultModel struct {
	Date  *string `json:"date,omitempty"`
	Count *int    `json:"count,omitempty"`
}

// DailyGroupEventCountRequestModel represents a request for daily group event counts.
type DailyGroupEventCountRequestModel struct {
	GroupID *int64  `json:"groupId,omitempty"`
	Start   *string `json:"start,omitempty"`
	End     *string `json:"end,omitempty"`
}

// CheckEventConflictInput represents input for checking event time conflicts.
type CheckEventConflictInput struct {
	Start                 *string `json:"start,omitempty"`
	End                   *string `json:"end,omitempty"`
	AllDay                bool    `json:"allDay"`
	InstitutionProfileIDs []int64 `json:"institutionProfileIds,omitempty"`
	ExcludeEventID        *int64  `json:"excludeEventId,omitempty"`
}

// ConflictEventItem represents a profile that conflicts with a proposed event time.
type ConflictEventItem struct {
	ProfileID *int64  `json:"profileId,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// CalendarResourceConflict represents a resource conflict result.
type CalendarResourceConflict struct {
	UnavailableResourceIDs []int64 `json:"unavailableResourceIds,omitempty"`
}

// CommunicationEventSpinnerItem represents a communication event spinner item.
type CommunicationEventSpinnerItem struct {
	DisplayName *string         `json:"displayName,omitempty"`
	Value       json.RawMessage `json:"value,omitempty"`
}

// ConversationMeetingUpdateInput represents an update input for a conversation meeting.
type ConversationMeetingUpdateInput struct {
	EventID *int64 `json:"eventId,omitempty"`
	ChildID *int64 `json:"childId,omitempty"`
}

// CreateConversationMeetingProcessedResult represents a result after creating a conversation meeting.
type CreateConversationMeetingProcessedResult struct {
	EventID          *int64                    `json:"eventId,omitempty"`
	ResourceConflict *CalendarResourceConflict `json:"resourceConflict,omitempty"`
}

// InviteeGroupRequest represents a group invitation request.
type InviteeGroupRequest struct {
	GroupID     *int64   `json:"groupId,omitempty"`
	PortalRoles []string `json:"portalRoles,omitempty"`
}

// CreateBaseEventRequest represents base fields for event creation requests.
type CreateBaseEventRequest struct {
	EventID              *int64                `json:"eventId,omitempty"`
	FromInstProfileID    *int64                `json:"fromInstProfileId,omitempty"`
	Title                *string               `json:"title,omitempty"`
	EventTypeEnum        *string               `json:"eventTypeEnum,omitempty"`
	Description          *string               `json:"description,omitempty"`
	InviteeIDs           []int64               `json:"inviteeIds,omitempty"`
	InviteeGroups        []InviteeGroupRequest `json:"inviteeGroups,omitempty"`
	InvitedGroupIDs      []int64               `json:"invitedGroupIds,omitempty"`
	CoOrganizerIDs       []int64               `json:"coOrganizerIds,omitempty"`
	InvitedOtpInboxIDs   []int64               `json:"invitedOtpInboxIds,omitempty"`
	AttachmentIDs        []int64               `json:"attachmentIds,omitempty"`
	HideInOwnCalendar    bool                  `json:"hideInOwnCalendar"`
	ResponseDeadline     *string               `json:"responseDeadline,omitempty"`
	InstitutionCode      *string               `json:"institutionCode,omitempty"`
}

// CreateSimpleEventRequest represents creating a simple event.
type CreateSimpleEventRequest struct {
	EventID                       *int64                `json:"eventId,omitempty"`
	FromInstProfileID             *int64                `json:"fromInstProfileId,omitempty"`
	Title                         *string               `json:"title,omitempty"`
	EventTypeEnum                 *string               `json:"eventTypeEnum,omitempty"`
	Description                   *string               `json:"description,omitempty"`
	InviteeIDs                    []int64               `json:"inviteeIds,omitempty"`
	InviteeGroups                 []InviteeGroupRequest `json:"inviteeGroups,omitempty"`
	InvitedGroupIDs               []int64               `json:"invitedGroupIds,omitempty"`
	CoOrganizerIDs                []int64               `json:"coOrganizerIds,omitempty"`
	InvitedOtpInboxIDs            []int64               `json:"invitedOtpInboxIds,omitempty"`
	AttachmentIDs                 []int64               `json:"attachmentIds,omitempty"`
	HideInOwnCalendar             bool                  `json:"hideInOwnCalendar"`
	ResponseDeadline              *string               `json:"responseDeadline,omitempty"`
	InstitutionCode               *string               `json:"institutionCode,omitempty"`
	StartDateTime                 *string               `json:"startDateTime,omitempty"`
	EndDateTime                   *string               `json:"endDateTime,omitempty"`
	AllDay                        bool                  `json:"allDay"`
	Private                       bool                  `json:"private"`
	ResponseRequired              bool                  `json:"responseRequired"`
	PrimaryResourceID             *int64                `json:"primaryResourceId,omitempty"`
	PrimaryResourceText           *string               `json:"primaryResourceText,omitempty"`
	AdditionalResourceIDs         []int64               `json:"additionalResourceIds,omitempty"`
	AdditionalResourceText        *string               `json:"additionalResourceText,omitempty"`
	AddToInstitutionCalendar      bool                  `json:"addToInstitutionCalendar"`
	AddedToInstitutionCalendar    bool                  `json:"addedToInstitutionCalendar"`
	MaximumNumberOfParticipants   *int64                `json:"maximumNumberOfParticipants,omitempty"`
	DoRequestNumberOfParticipants bool                  `json:"doRequestNumberOfParticipants"`
}

// CreateRepeatingEventRequest represents creating a repeating event.
type CreateRepeatingEventRequest struct {
	EventID                       *int64                `json:"eventId,omitempty"`
	FromInstProfileID             *int64                `json:"fromInstProfileId,omitempty"`
	Title                         *string               `json:"title,omitempty"`
	EventTypeEnum                 *string               `json:"eventTypeEnum,omitempty"`
	Description                   *string               `json:"description,omitempty"`
	InviteeIDs                    []int64               `json:"inviteeIds,omitempty"`
	InviteeGroups                 []InviteeGroupRequest `json:"inviteeGroups,omitempty"`
	InvitedGroupIDs               []int64               `json:"invitedGroupIds,omitempty"`
	CoOrganizerIDs                []int64               `json:"coOrganizerIds,omitempty"`
	InvitedOtpInboxIDs            []int64               `json:"invitedOtpInboxIds,omitempty"`
	AttachmentIDs                 []int64               `json:"attachmentIds,omitempty"`
	HideInOwnCalendar             bool                  `json:"hideInOwnCalendar"`
	ResponseDeadline              *string               `json:"responseDeadline,omitempty"`
	InstitutionCode               *string               `json:"institutionCode,omitempty"`
	StartDateTime                 *string               `json:"startDateTime,omitempty"`
	EndDateTime                   *string               `json:"endDateTime,omitempty"`
	AllDay                        bool                  `json:"allDay"`
	Private                       bool                  `json:"private"`
	ResponseRequired              bool                  `json:"responseRequired"`
	PrimaryResourceID             *int64                `json:"primaryResourceId,omitempty"`
	PrimaryResourceText           *string               `json:"primaryResourceText,omitempty"`
	AdditionalResourceIDs         []int64               `json:"additionalResourceIds,omitempty"`
	AdditionalResourceText        *string               `json:"additionalResourceText,omitempty"`
	AddToInstitutionCalendar      bool                  `json:"addToInstitutionCalendar"`
	AddedToInstitutionCalendar    bool                  `json:"addedToInstitutionCalendar"`
	MaximumNumberOfParticipants   *int64                `json:"maximumNumberOfParticipants,omitempty"`
	DoRequestNumberOfParticipants bool                  `json:"doRequestNumberOfParticipants"`
	OccurenceLimit                *int                  `json:"occurenceLimit,omitempty"`
	WeekdayMask                   []bool                `json:"weekdayMask,omitempty"`
	DayInMonth                    *int                  `json:"dayInMonth,omitempty"`
	RepeatTypeEnum                *string               `json:"repeatTypeEnum,omitempty"`
	Interval                      *int                  `json:"interval,omitempty"`
	MaxDate                       *string               `json:"maxDate,omitempty"`
	OccurrenceDateTime            *string               `json:"occurrenceDateTime,omitempty"`
}

// CreateTimeslotEventTimeSlotDto represents a timeslot DTO for timeslot event creation.
type CreateTimeslotEventTimeSlotDto struct {
	ID                  *int64  `json:"id,omitempty"`
	PrimaryResourceID   *int64  `json:"primaryResourceId,omitempty"`
	PrimaryResourceText *string `json:"primaryResourceText,omitempty"`
	StartDate           *string `json:"startDate,omitempty"`
	EndDate             *string `json:"endDate,omitempty"`
}

// CreateTimeslotEventRequest represents creating a timeslot event.
type CreateTimeslotEventRequest struct {
	EventID                         *int64                           `json:"eventId,omitempty"`
	FromInstProfileID               *int64                           `json:"fromInstProfileId,omitempty"`
	Title                           *string                          `json:"title,omitempty"`
	EventTypeEnum                   *string                          `json:"eventTypeEnum,omitempty"`
	Description                     *string                          `json:"description,omitempty"`
	InviteeIDs                      []int64                          `json:"inviteeIds,omitempty"`
	InviteeGroups                   []InviteeGroupRequest            `json:"inviteeGroups,omitempty"`
	InvitedGroupIDs                 []int64                          `json:"invitedGroupIds,omitempty"`
	CoOrganizerIDs                  []int64                          `json:"coOrganizerIds,omitempty"`
	InvitedOtpInboxIDs              []int64                          `json:"invitedOtpInboxIds,omitempty"`
	AttachmentIDs                   []int64                          `json:"attachmentIds,omitempty"`
	HideInOwnCalendar               bool                             `json:"hideInOwnCalendar"`
	ResponseDeadline                *string                          `json:"responseDeadline,omitempty"`
	InstitutionCode                 *string                          `json:"institutionCode,omitempty"`
	TimeSlots                       []CreateTimeslotEventTimeSlotDto `json:"timeSlots,omitempty"`
	BreakLength                     *int                             `json:"breakLength,omitempty"`
	MeetingDuration                 *int                             `json:"meetingDuration,omitempty"`
	ChildRequired                   bool                             `json:"childRequired"`
	MeetingsBetweenBreaks           *int                             `json:"meetingsBetweenBreaks,omitempty"`
	AddToInstitutionCalendar        bool                             `json:"addToInstitutionCalendar"`
	NumberOfParticipantsPerTimeSlot *int                             `json:"numberOfParticipantsPerTimeSlot,omitempty"`
}

// CreateEventResource represents a resource for event creation.
type CreateEventResource struct {
	ID                           *int64  `json:"id,omitempty"`
	InstitutionCode              *string `json:"institutionCode,omitempty"`
	InstitutionName              *string `json:"institutionName,omitempty"`
	Description                  *string `json:"description,omitempty"`
	Name                         *string `json:"name,omitempty"`
	ShortName                    *string `json:"shortName,omitempty"`
	DisplayName                  *string `json:"displayName,omitempty"`
	ResourceCategory             *string `json:"resourceCategory,omitempty"`
	NumberOfAvailableOccurrences *int    `json:"numberOfAvailableOccurrences,omitempty"`
	NumberOfOccurrences          *int    `json:"numberOfOccurrences,omitempty"`
}

// CreateEventResourceProfile represents a profile for event resource creation.
type CreateEventResourceProfile struct {
	ProfileID       *int64  `json:"profileId,omitempty"`
	Name            *string `json:"name,omitempty"`
	InstitutionCode *string `json:"institutionCode,omitempty"`
}

// UpdateLessonRequest represents updating a lesson.
type UpdateLessonRequest struct {
	EventID                *int64  `json:"eventId,omitempty"`
	InstitutionProfileID   *int64  `json:"institutionProfileId,omitempty"`
	NoteToClass            *string `json:"noteToClass,omitempty"`
	NoteToTeacher          *string `json:"noteToTeacher,omitempty"`
	NoteToSubstitute       *string `json:"noteToSubstitute,omitempty"`
	AdditionalResourceIDs  []int64 `json:"additionalResourceIds,omitempty"`
	AdditionalResourceText *string `json:"additionalResourceText,omitempty"`
	AttachmentIDs          []int64 `json:"attachmentIds,omitempty"`
}

// RespondSimpleEventRequest represents responding to a simple event invitation.
type RespondSimpleEventRequest struct {
	EventID                    *int64  `json:"eventId,omitempty"`
	InstitutionProfileID       *int64  `json:"institutionProfileId,omitempty"`
	InvitedInstProfileID       *int64  `json:"invitedInstProfileId,omitempty"`
	ResponseType               *string `json:"responseType,omitempty"`
	OccurrenceDateTime         *string `json:"occurrenceDateTime,omitempty"`
	NumberOfAdultParticipants  *int    `json:"numberOfAdultParticipants,omitempty"`
	NumberOfChildParticipants  *int    `json:"numberOfChildParticipants,omitempty"`
}

// RespondTimeslotEventRequest represents responding to a timeslot event.
type RespondTimeslotEventRequest struct {
	EventID                 *int64  `json:"eventId,omitempty"`
	ResponseTypeEnum        *string `json:"responseTypeEnum,omitempty"`
	TimeSlotID              *int64  `json:"timeSlotId,omitempty"`
	TimeSlotIndex           *int    `json:"timeSlotIndex,omitempty"`
	InstitutionProfileID    *int64  `json:"institutionProfileId,omitempty"`
	ConcerningInstProfileID *int64  `json:"concerningInstProfileId,omitempty"`
	OnBehalfOf              bool    `json:"onBehalfOf"`
}

// BlockTimeSlotRequest represents blocking a timeslot.
type BlockTimeSlotRequest struct {
	EventID       *int64 `json:"eventId,omitempty"`
	TimeSlotID    *int64 `json:"timeSlotId,omitempty"`
	TimeSlotIndex *int   `json:"timeSlotIndex,omitempty"`
}

// DeleteTimeslotRequest represents deleting a timeslot booking.
type DeleteTimeslotRequest struct {
	EventID                        *int64 `json:"eventId,omitempty"`
	TimeSlotID                     *int64 `json:"timeSlotId,omitempty"`
	TimeSlotIndex                  *int   `json:"timeSlotIndex,omitempty"`
	ConcerningInstitutionProfileID *int   `json:"concerningInstitutionProfileId,omitempty"`
}

// SendEventReminderRequest represents sending an event reminder.
type SendEventReminderRequest struct {
	EventID *int64  `json:"eventId,omitempty"`
	Message *string `json:"message,omitempty"`
}

// RelationsMessage represents a relations message for calendar event views.
type RelationsMessage struct {
	ProfileID    *int64  `json:"profileId,omitempty"`
	IsSelected   bool    `json:"isSelected"`
	RelationMode *string `json:"relationMode,omitempty"`
}

// GetEventsParameters represents parameters for getting events.
type GetEventsParameters struct {
	InstProfileIDs                  []int64  `json:"instProfileIds,omitempty"`
	ResourceIDs                     []int64  `json:"resourceIds,omitempty"`
	Start                           *string  `json:"start,omitempty"`
	End                             *string  `json:"end,omitempty"`
	SpecificTypes                   []string `json:"specificTypes,omitempty"`
	SchoolCalendarInstitutionCodes  []string `json:"schoolCalendarInstitutionCodes,omitempty"`
}

// GetEventsForInstitutionRequestModel represents parameters for getting institution events.
type GetEventsForInstitutionRequestModel struct {
	Start     *string  `json:"start,omitempty"`
	End       *string  `json:"end,omitempty"`
	InstCodes []string `json:"instCodes,omitempty"`
}

// GetEventTypeRequestDto represents a request to get event types for filtering.
type GetEventTypeRequestDto struct {
	FilterInstitutionCodes []string `json:"filterInstitutionCodes,omitempty"`
}

// VacationRegistrationChildrenCountByDates represents children count by date for vacation.
type VacationRegistrationChildrenCountByDates struct {
	Date              *string `json:"date,omitempty"`
	ChildrenAreComing *int    `json:"childrenAreComing,omitempty"`
	Total             *int    `json:"total,omitempty"`
}

// VacationRegistrationDetailsResultDto represents vacation registration details.
type VacationRegistrationDetailsResultDto struct {
	ChildrenTotal                *int                                        `json:"childrenTotal,omitempty"`
	ChildrenPendingAnswers       []ChildMetadata                             `json:"childrenPendingAnswers,omitempty"`
	VacationChildrenCountByDates []VacationRegistrationChildrenCountByDates  `json:"vacationChildrenCountByDates,omitempty"`
	NoteToGuardians              *string                                     `json:"noteToGuardians,omitempty"`
	Departments                  []json.RawMessage                           `json:"departments,omitempty"`
}

// VacationChildrenDto represents child info in a vacation day.
type VacationChildrenDto struct {
	Child                           *ChildMetadata `json:"child,omitempty"`
	Status                          *string        `json:"status,omitempty"`
	VacationRegistrationResponseID  *int64         `json:"vacationRegistrationResponseId,omitempty"`
}

// VacationDayDto represents a single vacation day with children responses.
type VacationDayDto struct {
	Date     *string               `json:"date,omitempty"`
	Children []VacationChildrenDto `json:"children,omitempty"`
}

// VacationWeekResultDto represents a vacation week overview.
type VacationWeekResultDto struct {
	FromDate      *string          `json:"fromDate,omitempty"`
	ToDate        *string          `json:"toDate,omitempty"`
	WeekNumber    *int             `json:"weekNumber,omitempty"`
	VacationDays  []VacationDayDto `json:"vacationDays,omitempty"`
}

// VacationOverviewListItemResultDto represents a vacation overview list item.
type VacationOverviewListItemResultDto struct {
	Title           *string `json:"title,omitempty"`
	ID              *int64  `json:"id,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
}

// VacationOverviewListRequestDto represents a request to list vacation overviews.
type VacationOverviewListRequestDto struct {
	FilterInstitutionCalendarCodes []string `json:"filterInstitutionCalendarCodes,omitempty"`
}

// CheckVacationRequestAnsweredRequestModel represents checking if vacation request has been answered.
type CheckVacationRequestAnsweredRequestModel struct {
	VacationRegistrationResponseID *int64 `json:"vacationRegistrationResponseId,omitempty"`
}

// GetVacationRequestResponseRequestModel represents getting vacation request response details.
type GetVacationRequestResponseRequestModel struct {
	VacationRequestID                     *int64  `json:"vacationRequestId,omitempty"`
	FilterDepartmentGroupIDs              []int64 `json:"filterDepartmentGroupIds,omitempty"`
	FilterDepartmentFilteringGroupIDs     []int64 `json:"filterDepartmentFilteringGroupIds,omitempty"`
}

// RespondToVacationRegistrationRequestDto represents responding to a vacation registration request.
type RespondToVacationRegistrationRequestDto struct {
	ChildID                        *int64                              `json:"childId,omitempty"`
	VacationRegistrationResponseID *int64                              `json:"vacationRegistrationResponseId,omitempty"`
	Days                           []CalendarGuardianRegisterVacationIntervals `json:"days,omitempty"`
	Comment                        *string                             `json:"comment,omitempty"`
}

// CalendarGuardianRegisterVacationIntervals represents vacation day intervals (guardian registration) in calendar context.
type CalendarGuardianRegisterVacationIntervals struct {
	Date      *string `json:"date,omitempty"`
	EntryTime *string `json:"entryTime,omitempty"`
	ExitTime  *string `json:"exitTime,omitempty"`
	IsComing  bool    `json:"isComing"`
}

// VacationDetailsDto represents vacation details DTO.
type VacationDetailsDto struct {
	IsVacationCreatedFromVacationRequest bool `json:"isVacationCreatedFromVacationRequest"`
}

// SimpleGroupWithRolesModel represents a simple group with portal roles.
type SimpleGroupWithRolesModel struct {
	ID          *int64   `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	PortalRoles []string `json:"portalRoles,omitempty"`
}
