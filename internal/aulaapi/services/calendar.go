package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GetEvents gets calendar events with filtering by profile, date range, and type.
func GetEvents(ctx context.Context, s *aulaapi.Session, params *models.GetEventsParameters) ([]models.EventSimpleDto, error) {
	return aulaapi.SessionPost[[]models.EventSimpleDto](ctx, s, "?method=calendar.getEventsByProfileIdsAndResourceIds", params)
}

// GetEventDetail gets event detail by ID.
func GetEventDetail(ctx context.Context, s *aulaapi.Session, eventID int64) (models.EventDetailsDto, error) {
	return aulaapi.SessionGet[models.EventDetailsDto](ctx, s, fmt.Sprintf("?method=calendar.getEventById&eventId=%d", eventID))
}

// GetDailyAggregatedEvents gets daily aggregated events (event counts per day per type).
func GetDailyAggregatedEvents(ctx context.Context, s *aulaapi.Session, params *models.GetEventsParameters) ([]models.DailyAggregatedEventsResultModel, error) {
	var query []string
	if params.InstProfileIDs != nil {
		for _, id := range params.InstProfileIDs {
			query = append(query, ParamNum("instProfileIds", id))
		}
	}
	if params.Start != nil {
		query = append(query, fmt.Sprintf("start=%s", EncodeValue(*params.Start)))
	}
	if params.End != nil {
		query = append(query, fmt.Sprintf("end=%s", EncodeValue(*params.End)))
	}
	if params.SpecificTypes != nil {
		for _, t := range params.SpecificTypes {
			query = append(query, fmt.Sprintf("specificTypes=%s", EncodeValue(t)))
		}
	}
	if params.SchoolCalendarInstitutionCodes != nil {
		for _, c := range params.SchoolCalendarInstitutionCodes {
			query = append(query, fmt.Sprintf("schoolCalendarInstitutionCodes=%s", EncodeValue(c)))
		}
	}
	path := "?method=calendar.getDailyAggregatedEvents"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.DailyAggregatedEventsResultModel](ctx, s, path)
}

// GetDailyGroupEventCount gets daily event count per group.
func GetDailyGroupEventCount(ctx context.Context, s *aulaapi.Session, groupID int64, start, end string) ([]models.DailyEventCountResultModel, error) {
	path := fmt.Sprintf("?method=calendar.getDailyEventCountForGroup&groupId=%d&start=%s&end=%s", groupID, EncodeValue(start), EncodeValue(end))
	return aulaapi.SessionGet[[]models.DailyEventCountResultModel](ctx, s, path)
}

// GetEventForGroup gets events for a specific group.
func GetEventForGroup(ctx context.Context, s *aulaapi.Session, groupID int64, start, end *string) ([]models.EventSimpleDto, error) {
	var query []string
	if start != nil {
		query = append(query, fmt.Sprintf("start=%s", EncodeValue(*start)))
	}
	if end != nil {
		query = append(query, fmt.Sprintf("end=%s", EncodeValue(*end)))
	}
	path := fmt.Sprintf("?method=calendar.geteventsbygroupid&groupId=%d", groupID)
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.EventSimpleDto](ctx, s, path)
}

// GetSchoolEvents gets school-wide events.
func GetSchoolEvents(ctx context.Context, s *aulaapi.Session, params *models.GetEventsForInstitutionRequestModel) ([]models.EventSimpleDto, error) {
	var query []string
	if params.Start != nil {
		query = append(query, fmt.Sprintf("start=%s", EncodeValue(*params.Start)))
	}
	if params.End != nil {
		query = append(query, fmt.Sprintf("end=%s", EncodeValue(*params.End)))
	}
	if params.InstCodes != nil {
		for _, c := range params.InstCodes {
			query = append(query, fmt.Sprintf("instCodes=%s", EncodeValue(c)))
		}
	}
	path := "?method=calendar.getEventsForInstitutions"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.EventSimpleDto](ctx, s, path)
}

// GetEventTypes gets available event types for filtering.
func GetEventTypes(ctx context.Context, s *aulaapi.Session, filterInstitutionCodes []string) (models.GetEventTypesByPortalRoleResultModel, error) {
	var query []string
	for _, code := range filterInstitutionCodes {
		query = append(query, fmt.Sprintf("filterInstitutionCodes=%s", EncodeValue(code)))
	}
	path := "?method=calendar.getEventTypes"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetEventTypesByPortalRoleResultModel](ctx, s, path)
}

// GetEventTypesForCalendarFeed gets event types available for calendar feed configuration.
func GetEventTypesForCalendarFeed(ctx context.Context, s *aulaapi.Session) (models.GetEventTypesByPortalRoleResultModel, error) {
	return aulaapi.SessionGet[models.GetEventTypesByPortalRoleResultModel](ctx, s, "?method=CalendarFeed.getEventTypesRelevantForPortalRole")
}

// DeleteEvent deletes a calendar event.
func DeleteEvent(ctx context.Context, s *aulaapi.Session, eventID int64) (json.RawMessage, error) {
	body := map[string]int64{"eventId": eventID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.deleteEvent", body)
}

// RespondSimpleEvent responds to a simple event invitation (accept/decline).
func RespondSimpleEvent(ctx context.Context, s *aulaapi.Session, args *models.RespondSimpleEventRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.respondToSimpleEvent", args)
}

// RespondTimeslotEvent responds to a timeslot event (book a timeslot).
func RespondTimeslotEvent(ctx context.Context, s *aulaapi.Session, args *models.RespondTimeslotEventRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.respondToTimeSlotEvent", args)
}

// EditTimeslotEvent edits a timeslot event.
func EditTimeslotEvent(ctx context.Context, s *aulaapi.Session, args *models.CreateTimeslotEventRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.updateResponseToTimeSlotEvent", args)
}

// BlockTimeSlot blocks a timeslot to prevent booking.
func BlockTimeSlot(ctx context.Context, s *aulaapi.Session, args *models.BlockTimeSlotRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.blockTimeSlot", args)
}

// DeleteTimeSlot deletes a timeslot booking.
func DeleteTimeSlot(ctx context.Context, s *aulaapi.Session, args *models.DeleteTimeslotRequest) (json.RawMessage, error) {
	eventID := int64(0)
	if args.EventID != nil {
		eventID = *args.EventID
	}
	var query []string
	if args.TimeSlotID != nil {
		query = append(query, fmt.Sprintf("timeSlotId=%d", *args.TimeSlotID))
	}
	if args.TimeSlotIndex != nil {
		query = append(query, fmt.Sprintf("timeSlotIndex=%d", *args.TimeSlotIndex))
	}
	if args.ConcerningInstitutionProfileID != nil {
		query = append(query, fmt.Sprintf("concerningInstitutionProfileId=%d", *args.ConcerningInstitutionProfileID))
	}
	path := fmt.Sprintf("?method=calendar.removeBlockingOrResponseToTimeSlot&eventId=%d", eventID)
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, path)
}

// UpdateLessonEvent updates a lesson event (notes, resources, attachments).
func UpdateLessonEvent(ctx context.Context, s *aulaapi.Session, args *models.UpdateLessonRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.updateLessonEvent", args)
}

// AddVacation adds a vacation registration event.
func AddVacation(ctx context.Context, s *aulaapi.Session, args *models.CreateSimpleEventRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.addVacation", args)
}

// GetVacation gets a vacation by ID.
func GetVacation(ctx context.Context, s *aulaapi.Session, vacationID int64) (models.EventDetailsDto, error) {
	return aulaapi.SessionGet[models.EventDetailsDto](ctx, s, fmt.Sprintf("?method=calendar.getVacationById&vacationId=%d", vacationID))
}

// DeleteVacation deletes a vacation.
func DeleteVacation(ctx context.Context, s *aulaapi.Session, vacationID int64) (json.RawMessage, error) {
	body := map[string]int64{"vacationId": vacationID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.deleteVacation", body)
}

// GetFutureVacationRequest gets future vacation requests.
func GetFutureVacationRequest(ctx context.Context, s *aulaapi.Session, filterInstitutionCodes []string) ([]models.VacationOverviewListItemResultDto, error) {
	var query []string
	for _, code := range filterInstitutionCodes {
		query = append(query, fmt.Sprintf("filterInstitutionCalendarCodes=%s", EncodeValue(code)))
	}
	path := "?method=calendar.getFutureVacationRequests"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.VacationOverviewListItemResultDto](ctx, s, path)
}

// GetVacationRequestResponse gets the response details for a vacation request.
func GetVacationRequestResponse(ctx context.Context, s *aulaapi.Session, args *models.GetVacationRequestResponseRequestModel) ([]models.VacationWeekResultDto, error) {
	vacationID := int64(0)
	if args.VacationRequestID != nil {
		vacationID = *args.VacationRequestID
	}
	var query []string
	if args.FilterDepartmentGroupIDs != nil {
		for _, id := range args.FilterDepartmentGroupIDs {
			query = append(query, ParamNum("filterDepartmentGroupIds", id))
		}
	}
	if args.FilterDepartmentFilteringGroupIDs != nil {
		for _, id := range args.FilterDepartmentFilteringGroupIDs {
			query = append(query, ParamNum("filterDepartmentFilteringGroupIds", id))
		}
	}
	path := fmt.Sprintf("?method=calendar.getVacationRequestResponses&vacationRequestId=%d", vacationID)
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.VacationWeekResultDto](ctx, s, path)
}

// RespondToVacationRegistrationRequest responds to a vacation registration request.
func RespondToVacationRegistrationRequest(ctx context.Context, s *aulaapi.Session, args *models.RespondToVacationRegistrationRequestDto) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.respondToVacationRegistrationRequest", args)
}

// GetCalendarSynchronisationConfigurations gets all calendar synchronisation configurations.
func GetCalendarSynchronisationConfigurations(ctx context.Context, s *aulaapi.Session) ([]models.CalendarSynchronisationConfigurationItem, error) {
	return aulaapi.SessionGet[[]models.CalendarSynchronisationConfigurationItem](ctx, s, "?method=CalendarFeed.getFeedConfigurations")
}

// CreateCalendarSynchronisationConfiguration creates a new calendar synchronisation configuration.
func CreateCalendarSynchronisationConfiguration(ctx context.Context, s *aulaapi.Session, args *models.CreateCalendarSynchronizationConfigurationRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=CalendarFeed.createFeedConfiguration", args)
}

// UpdateCalendarSynchronisationConfiguration updates an existing calendar synchronisation configuration.
func UpdateCalendarSynchronisationConfiguration(ctx context.Context, s *aulaapi.Session, args *models.UpdateCalendarSynchronizationConfigurationRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=CalendarFeed.updateFeedConfiguration", args)
}

// DeleteCalendarSynchronisationConfiguration deletes a calendar synchronisation configuration.
func DeleteCalendarSynchronisationConfiguration(ctx context.Context, s *aulaapi.Session, configID int64) (json.RawMessage, error) {
	body := map[string]int64{"configId": configID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=CalendarFeed.removeFeedConfiguration", body)
}

// GetCalendarSynchronisationConsent gets the current calendar synchronisation consent status.
func GetCalendarSynchronisationConsent(ctx context.Context, s *aulaapi.Session) (models.CalendarSynchronisationModel, error) {
	return aulaapi.SessionGet[models.CalendarSynchronisationModel](ctx, s, "?method=CalendarFeed.getPolicyAnswer")
}

// UpdateCalendarSynchronisationConsent updates (accept/revoke) calendar synchronisation consent.
func UpdateCalendarSynchronisationConsent(ctx context.Context, s *aulaapi.Session, args *models.CalendarSynchronisationModel) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=CalendarFeed.setPolicyAnswer", args)
}

// GetDelegatedAccesses gets delegated calendar accesses for the current user.
func GetDelegatedAccesses(ctx context.Context, s *aulaapi.Session, instProfileID *int64) ([]models.DelegateAccesses, error) {
	path := "?method=calendar.getDelegatedAccesses"
	if instProfileID != nil {
		path = fmt.Sprintf("?method=calendar.getDelegatedAccesses&instProfileId=%d", *instProfileID)
	}
	return aulaapi.SessionGet[[]models.DelegateAccesses](ctx, s, path)
}

// SetDelegatedAccesses sets delegated calendar accesses.
func SetDelegatedAccesses(ctx context.Context, s *aulaapi.Session, args *models.DelegateAccessesInput) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.setDelegatedAccesses", args)
}

// GetInstitutionProfilesWithDelegatedAccesses gets institution profiles that have delegated calendar access.
func GetInstitutionProfilesWithDelegatedAccesses(ctx context.Context, s *aulaapi.Session, instProfileID *int64) ([]models.InstitutionDelegateAccessesItem, error) {
	path := "?method=calendar.getInstitutionProfilesWithDelegatedAccess"
	if instProfileID != nil {
		path = fmt.Sprintf("?method=calendar.getInstitutionProfilesWithDelegatedAccess&instProfileId=%d", *instProfileID)
	}
	return aulaapi.SessionGet[[]models.InstitutionDelegateAccessesItem](ctx, s, path)
}

// GetBirthdaysForGroup gets birthdays for a group.
func GetBirthdaysForGroup(ctx context.Context, s *aulaapi.Session, groupID int64, start, end string) ([]models.BirthdayEventDto, error) {
	path := fmt.Sprintf("?method=calendar.getBirthdayEventsForGroup&groupId=%d&start=%s&end=%s", groupID, EncodeValue(start), EncodeValue(end))
	return aulaapi.SessionGet[[]models.BirthdayEventDto](ctx, s, path)
}

// GetBirthdaysForInstitution gets birthdays for an institution.
func GetBirthdaysForInstitution(ctx context.Context, s *aulaapi.Session, institutionID int64, start, end string) ([]models.BirthdayEventDto, error) {
	path := fmt.Sprintf("?method=calendar.getBirthdayEventsForInstitutions&institutionId=%d&start=%s&end=%s", institutionID, EncodeValue(start), EncodeValue(end))
	return aulaapi.SessionGet[[]models.BirthdayEventDto](ctx, s, path)
}

// GetTopImportantDate gets the top important dates (shown on dashboard).
func GetTopImportantDate(ctx context.Context, s *aulaapi.Session, instProfileIDs []int64) ([]models.ImportantDateItem, error) {
	var query []string
	for _, id := range instProfileIDs {
		query = append(query, ParamNum("instProfileIds", id))
	}
	path := "?method=calendar.getImportantDates"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.ImportantDateItem](ctx, s, path)
}

// CheckConflictEventForAttendees checks scheduling conflicts for attendees.
func CheckConflictEventForAttendees(ctx context.Context, s *aulaapi.Session, args *models.CheckEventConflictInput) ([]models.ConflictEventItem, error) {
	return aulaapi.SessionPost[[]models.ConflictEventItem](ctx, s, "?method=calendar.checkConflictEventForAttendees", args)
}

// GetIsCalendarFeedEnabledForMunicipality checks whether calendar feed is enabled for a municipality.
func GetIsCalendarFeedEnabledForMunicipality(ctx context.Context, s *aulaapi.Session, municipalityID int64) (models.CalendarSynchronisationMunicipalityFeedModel, error) {
	return aulaapi.SessionGet[models.CalendarSynchronisationMunicipalityFeedModel](ctx, s, fmt.Sprintf("?method=MunicipalConfiguration.getCalendarFeedEnabled&municipalityId=%d", municipalityID))
}

// GetFeedConfigurationByID gets a feed configuration by ID.
func GetFeedConfigurationByID(ctx context.Context, s *aulaapi.Session, configID int64) (models.CalendarSynchronisationConfigurationItem, error) {
	return aulaapi.SessionGet[models.CalendarSynchronisationConfigurationItem](ctx, s, fmt.Sprintf("?method=CalendarFeed.getFeedConfigurationById&configId=%d", configID))
}
