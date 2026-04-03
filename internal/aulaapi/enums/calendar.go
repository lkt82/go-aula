package enums

// EventClass classifies a calendar event.
type EventClass string

const (
	EventClassBasic   EventClass = "basic"
	EventClassSeries  EventClass = "series"
	EventClassTimeslot EventClass = "timeslot"
	EventClassLesson  EventClass = "lesson"
	EventClassUnknown EventClass = "unknown"
)

// EventPlacementComparedToDateTime describes how an event is placed relative to a date/time.
type EventPlacementComparedToDateTime string

const (
	EventPlacementComparedToDateTimeNotOnTheDate              EventPlacementComparedToDateTime = "notOnTheDate"
	EventPlacementComparedToDateTimeStartAndEndOnDate         EventPlacementComparedToDateTime = "startAndEndOnDate"
	EventPlacementComparedToDateTimeStartOnDateButEndAfter    EventPlacementComparedToDateTime = "startOnDateButEndAfter"
	EventPlacementComparedToDateTimeStartBeforeDateButEndOn   EventPlacementComparedToDateTime = "startBeforeDateButEndOn"
	EventPlacementComparedToDateTimeStartBeforeAndEndAfterDate EventPlacementComparedToDateTime = "startBeforeAndEndAfterDate"
)

// EventPortraitType is the portrait display type for calendar events.
type EventPortraitType string

const (
	EventPortraitTypeEvent    EventPortraitType = "event"
	EventPortraitTypeBirthday EventPortraitType = "birthday"
	EventPortraitTypeAllDay   EventPortraitType = "allDay"
)

// EventType is the type of calendar event.
type EventType string

const (
	EventTypeEvent                EventType = "event"
	EventTypeHoliday              EventType = "holiday"
	EventTypePresenceHoliday      EventType = "presenceHoliday"
	EventTypeVacationRegistration EventType = "vacationRegistration"
	EventTypeBirthday             EventType = "birthday"
	EventTypeMeeting              EventType = "meeting"
	EventTypeOther                EventType = "other"
	EventTypeExcursion            EventType = "excursion"
	EventTypeSchoolHomeMeeting    EventType = "schoolHomeMeeting"
	EventTypeClassMeeting         EventType = "classMeeting"
	EventTypeParentalMeeting      EventType = "parentalMeeting"
	EventTypePerformanceMeeting   EventType = "performanceMeeting"
	EventTypeLesson               EventType = "lesson"
	EventTypeUnknown              EventType = "unknown"
)

// LessonStatus is the status of a lesson in the schedule.
type LessonStatus string

const (
	LessonStatusCancelled      LessonStatus = "cancelled"
	LessonStatusNormal         LessonStatus = "normal"
	LessonStatusAbsent         LessonStatus = "absent"
	LessonStatusSubstitute     LessonStatus = "substitute"
	LessonStatusToBeDeleted    LessonStatus = "toBeDeleted"
	LessonStatusWillBeUpdated  LessonStatus = "willBeUpdated"
	LessonStatusStatusNotFound LessonStatus = "statusNotFound"
)

// ParticipantRole is the role of a participant in a calendar event.
type ParticipantRole string

const (
	ParticipantRolePrimaryTeacher    ParticipantRole = "primaryTeacher"
	ParticipantRoleSubstituteTeacher ParticipantRole = "substituteTeacher"
	ParticipantRoleHelpTeacher       ParticipantRole = "helpTeacher"
	ParticipantRolePedagogue         ParticipantRole = "pedagogue"
	ParticipantRoleNotChosen         ParticipantRole = "notChosen"
)

// RepeatType describes how an event repeats.
type RepeatType string

const (
	RepeatTypeNever   RepeatType = "never"
	RepeatTypeDaily   RepeatType = "daily"
	RepeatTypeWeekly  RepeatType = "weekly"
	RepeatTypeMonthly RepeatType = "monthly"
)

// RepeatingEventDropdownEnum is a dropdown option when editing a repeating event.
type RepeatingEventDropdownEnum string

const (
	RepeatingEventDropdownEnumForSeries           RepeatingEventDropdownEnum = "forSeries"
	RepeatingEventDropdownEnumForSingleOccurrence RepeatingEventDropdownEnum = "forSingleOccurrence"
)

// ResponseType is the response to an event invitation.
type ResponseType string

const (
	ResponseTypeWaiting   ResponseType = "waiting"
	ResponseTypeDeclined  ResponseType = "declined"
	ResponseTypeAccepted  ResponseType = "accepted"
	ResponseTypeTentative ResponseType = "tentative"
)

// TimeslotResponseType is the availability status for a timeslot.
type TimeslotResponseType string

const (
	TimeslotResponseTypeBlocked       TimeslotResponseType = "blocked"
	TimeslotResponseTypeNotBooked     TimeslotResponseType = "notBooked"
	TimeslotResponseTypeAlreadyBooked TimeslotResponseType = "alreadyBooked"
)

// VacationRegistrationResponseStatus is the vacation registration response status.
type VacationRegistrationResponseStatus string

const (
	VacationRegistrationResponseStatusAnswered   VacationRegistrationResponseStatus = "answered"
	VacationRegistrationResponseStatusUnanswered VacationRegistrationResponseStatus = "unanswered"
)

// VacationResponseStatusEnum is the vacation response status.
type VacationResponseStatusEnum string

const (
	VacationResponseStatusEnumIsComing      VacationResponseStatusEnum = "isComing"
	VacationResponseStatusEnumIsNotComing   VacationResponseStatusEnum = "isNotComing"
	VacationResponseStatusEnumPendingAnswer VacationResponseStatusEnum = "pendingAnswer"
)

// RelationMode is the relation mode for calendar views.
type RelationMode string

const (
	RelationModeChildMode   RelationMode = "childMode"
	RelationModeInstitution RelationMode = "institution"
)

// CalendarItemType is the calendar item type.
type CalendarItemType string

const (
	CalendarItemTypeEvent    CalendarItemType = "event"
	CalendarItemTypeTitle    CalendarItemType = "title"
	CalendarItemTypeBirthday CalendarItemType = "birthday"
)

// MyCalendarItemType is the my-calendar item type.
type MyCalendarItemType string

const (
	MyCalendarItemTypeBody  MyCalendarItemType = "body"
	MyCalendarItemTypeTitle MyCalendarItemType = "title"
)
