package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// GetChildrensState gets children's current presence state.
func GetChildrensState(ctx context.Context, s *aulaapi.Session, instProfileIDs []int64) ([]models.ChildStatusDto, error) {
	var query []string
	for _, id := range instProfileIDs {
		query = append(query, fmt.Sprintf("institutionProfileIds[]=%d", id))
	}
	path := "?method=presence.getPresenceStates"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.ChildStatusDto](ctx, s, path)
}

// GetPresenceRegistrationDetail gets presence registration detail by ID.
func GetPresenceRegistrationDetail(ctx context.Context, s *aulaapi.Session, registrationID int64) (models.PresenceRegistrationResult, error) {
	return aulaapi.SessionGet[models.PresenceRegistrationResult](ctx, s, fmt.Sprintf("?method=presence.getPresenceRegistrationsByIds&registrationId=%d", registrationID))
}

// UpdatePresenceRegistration updates a presence registration (checkout details).
func UpdatePresenceRegistration(ctx context.Context, s *aulaapi.Session, args *models.UpdatePresenceRegistrationRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.updatePresenceRegistration", args)
}

// UpdateStatusByPresenceRegistrationIDs bulk updates presence status by registration IDs.
func UpdateStatusByPresenceRegistrationIDs(ctx context.Context, s *aulaapi.Session, args *models.BulkUpdatePresenceStatusRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.bulkUpdatePresenceStatus", args)
}

// UpdateStatusByInstitutionProfileIDs updates presence status by institution profile IDs.
func UpdateStatusByInstitutionProfileIDs(ctx context.Context, s *aulaapi.Session, args *models.UpdateStatusByInstitutionProfileIds) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.updateStatusByInstitutionProfileIds", args)
}

// GetPresenceSchedules gets presence schedules for children within a date range.
func GetPresenceSchedules(ctx context.Context, s *aulaapi.Session, args *models.PresenceSchedulesRequest) (json.RawMessage, error) {
	var query []string
	if args.FilterInstitutionProfileIDs != nil {
		for _, id := range args.FilterInstitutionProfileIDs {
			query = append(query, fmt.Sprintf("filterInstitutionProfileIds[]=%d", id))
		}
	}
	if args.FromDate != nil {
		query = append(query, fmt.Sprintf("fromDate=%s", EncodeValue(*args.FromDate)))
	}
	if args.ToDate != nil {
		query = append(query, fmt.Sprintf("toDate=%s", EncodeValue(*args.ToDate)))
	}
	path := "?method=presence.getPresenceTemplates"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[json.RawMessage](ctx, s, path)
}

// GetPresenceWeekOverview gets presence week overview (employee view).
func GetPresenceWeekOverview(ctx context.Context, s *aulaapi.Session, args *models.ComeGoGetWeekOverviewRequest) (models.GetPresenceOverview, error) {
	var query []string
	query = append(query, fmt.Sprintf("departmentId=%d", args.DepartmentID))
	if args.GroupIDs != nil {
		for _, id := range args.GroupIDs {
			query = append(query, fmt.Sprintf("groupIds=%d", id))
		}
	}
	if args.StatusFilters != nil {
		for _, f := range args.StatusFilters {
			query = append(query, fmt.Sprintf("statusFilters=%s", EncodeValue(f)))
		}
	}
	if args.StartDate != nil {
		query = append(query, fmt.Sprintf("startDate=%s", EncodeValue(*args.StartDate)))
	}
	if args.EndDate != nil {
		query = append(query, fmt.Sprintf("endDate=%s", EncodeValue(*args.EndDate)))
	}
	query = append(query, fmt.Sprintf("offset=%d", args.Offset))
	query = append(query, fmt.Sprintf("limit=%d", args.Limit))
	path := "?method=presence.getActivityOverview&" + strings.Join(query, "&")
	return aulaapi.SessionGet[models.GetPresenceOverview](ctx, s, path)
}

// UpdateOneDayPresence updates presence for a single day (template editing).
func UpdateOneDayPresence(ctx context.Context, s *aulaapi.Session, args *models.UpdatePresenceDayRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.updatePresenceTemplate", args)
}

// GetTemplateForDate gets the presence template for a specific date.
func GetTemplateForDate(ctx context.Context, s *aulaapi.Session, date string, institutionProfileID int64) (models.GetDayTemplateResult, error) {
	path := fmt.Sprintf("?method=presence.getTemplateForDate&date=%s&institutionProfileId=%d", date, institutionProfileID)
	return aulaapi.SessionGet[models.GetDayTemplateResult](ctx, s, path)
}

// DeleteRepeatedPresenceTemplate deletes a repeated presence template.
func DeleteRepeatedPresenceTemplate(ctx context.Context, s *aulaapi.Session, args *models.DeletePresenceTemplateRequest) (json.RawMessage, error) {
	templateID := int64(0)
	if args.PresentTemplateID != nil {
		templateID = *args.PresentTemplateID
	}
	var query []string
	if args.DeleteFromDay != nil {
		query = append(query, fmt.Sprintf("deleteFromDay=%s", EncodeValue(*args.DeleteFromDay)))
	}
	path := fmt.Sprintf("?method=presence.deleteRepeatingPresenceTemplate&templateId=%d", templateID)
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, path)
}

// GetOverlappingPresenceTemplates gets overlapping presence templates.
func GetOverlappingPresenceTemplates(ctx context.Context, s *aulaapi.Session, args *models.GetOverlappingPresenceTemplatesRequest) ([]json.RawMessage, error) {
	var query []string
	query = append(query, fmt.Sprintf("institutionProfileId=%d", args.InstitutionProfileID))
	if args.StartDate != nil {
		query = append(query, fmt.Sprintf("startDate=%s", EncodeValue(*args.StartDate)))
	}
	if args.EndDate != nil {
		query = append(query, fmt.Sprintf("endDate=%s", EncodeValue(*args.EndDate)))
	}
	if args.RepeatPattern != nil {
		query = append(query, fmt.Sprintf("repeatPattern=%s", *args.RepeatPattern))
	}
	path := "?method=presence.getOverlappingPresenceTemplates&" + strings.Join(query, "&")
	return aulaapi.SessionGet[[]json.RawMessage](ctx, s, path)
}

// GetSuggestionsForPickup gets pickup suggestions (exit-with suggestions).
func GetSuggestionsForPickup(ctx context.Context, s *aulaapi.Session, args *models.ComeGoExitWithSuggestionRequest) (models.GetExitWithSuggestionsResult, error) {
	var query []string
	if args.PickupName != nil {
		query = append(query, fmt.Sprintf("pickupName=%s", EncodeValue(*args.PickupName)))
	}
	if args.UniStudentIDs != nil {
		for _, id := range args.UniStudentIDs {
			query = append(query, fmt.Sprintf("uniStudentIds=%d", id))
		}
	}
	path := "?method=presence.getSuggestedNamesForPickupChild"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetExitWithSuggestionsResult](ctx, s, path)
}

// UpdateSuggestionsForPickup updates pickup suggestions (save pickup name).
func UpdateSuggestionsForPickup(ctx context.Context, s *aulaapi.Session, args *models.SavePickupNameRequest) (models.UpdatePickUpResponsibleResult, error) {
	return aulaapi.SessionPost[models.UpdatePickUpResponsibleResult](ctx, s, "?method=presence.savePickupNames", args)
}

// GetPickupResponsibles gets pickup responsibles for children.
func GetPickupResponsibles(ctx context.Context, s *aulaapi.Session, args *models.GetPickupResponsibleRequest) (models.GetPickupResponsibleResult, error) {
	var query []string
	if args.UniStudentIDs != nil {
		for _, id := range args.UniStudentIDs {
			query = append(query, fmt.Sprintf("uniStudentIds=%d", id))
		}
	}
	path := "?method=presence.getPickupResponsibles"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetPickupResponsibleResult](ctx, s, path)
}

// DeletePickupResponsible deletes a pickup responsible entry.
func DeletePickupResponsible(ctx context.Context, s *aulaapi.Session, args *models.DeletePickupResponsibleRequest) (json.RawMessage, error) {
	body := map[string]int64{"presencePickupSuggestionId": args.PresencePickupSuggestionID}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.deletePickupResponsible", body)
}

// GetChildGoHomeWith gets children that a child can go home with.
func GetChildGoHomeWith(ctx context.Context, s *aulaapi.Session, childID int64) (models.GetChildGoHomeWithResult, error) {
	return aulaapi.SessionGet[models.GetChildGoHomeWithResult](ctx, s, fmt.Sprintf("?method=presence.getGoHomeWithList&childId=%d", childID))
}

// AddSleepIntervals adds sleep intervals for children.
func AddSleepIntervals(ctx context.Context, s *aulaapi.Session, args *models.AddSleepIntervalsRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.addSleepIntervals", args)
}

// UpdateSleepInterval updates a sleep interval.
func UpdateSleepInterval(ctx context.Context, s *aulaapi.Session, args *models.UpdateSleepIntervalsDto) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.updateSleepInterval", args)
}

// DeleteSleepIntervals deletes sleep intervals.
func DeleteSleepIntervals(ctx context.Context, s *aulaapi.Session, sleepIntervalIDs []int64) (json.RawMessage, error) {
	body := map[string][]int64{"sleepIntervalIds": sleepIntervalIDs}
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.deleteSleepIntervals", body)
}

// GetActivityList gets the activity list for a department.
func GetActivityList(ctx context.Context, s *aulaapi.Session, args *models.ActivityListRequest) (models.ActivityListResult, error) {
	var query []string
	query = append(query, fmt.Sprintf("departmentId=%d", args.DepartmentID))
	if args.GroupIDs != nil {
		for _, id := range args.GroupIDs {
			query = append(query, fmt.Sprintf("groupIds=%d", id))
		}
	}
	if args.Limit != nil {
		query = append(query, fmt.Sprintf("limit=%d", *args.Limit))
	}
	if args.Offset != nil {
		query = append(query, fmt.Sprintf("offset=%d", *args.Offset))
	}
	if args.States != nil {
		for _, s := range args.States {
			query = append(query, fmt.Sprintf("states=%s", s))
		}
	}
	if args.NextActivity != nil {
		query = append(query, fmt.Sprintf("nextActivity=%s", *args.NextActivity))
	}
	if args.LocationIDs != nil {
		for _, id := range args.LocationIDs {
			query = append(query, fmt.Sprintf("locationIds=%d", id))
		}
	}
	if args.SortOn != nil {
		query = append(query, fmt.Sprintf("sortOn=%s", EncodeValue(*args.SortOn)))
	}
	path := "?method=presence.getActivityList&" + strings.Join(query, "&")
	return aulaapi.SessionGet[models.ActivityListResult](ctx, s, path)
}

// GetActivityFilter gets activity filter options for a department.
func GetActivityFilter(ctx context.Context, s *aulaapi.Session, institutionCode string) (models.ActivityFilterResult, error) {
	return aulaapi.SessionGet[models.ActivityFilterResult](ctx, s, fmt.Sprintf("?method=presence.getActivityListEditOptions&institutionCode=%s", institutionCode))
}

// GetDailyOverview gets daily presence overview (parent view).
func GetDailyOverview(ctx context.Context, s *aulaapi.Session, instProfileIDs []int64) ([]models.ParentsDailyOverviewResult, error) {
	var query []string
	for _, id := range instProfileIDs {
		query = append(query, fmt.Sprintf("childIds[]=%d", id))
	}
	path := "?method=presence.getDailyOverview"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.ParentsDailyOverviewResult](ctx, s, path)
}

// GetAvailableLocations gets available locations for presence tracking.
func GetAvailableLocations(ctx context.Context, s *aulaapi.Session, institutionCode string) ([]models.PresenceLocation, error) {
	return aulaapi.SessionGet[[]models.PresenceLocation](ctx, s, fmt.Sprintf("?method=presence.getAvailablePresenceLocations&institutionCode=%s", institutionCode))
}

// UpdateLocation updates location for children.
func UpdateLocation(ctx context.Context, s *aulaapi.Session, args *models.UpdateLocationRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=presence.updateLocation", args)
}

// PresenceAddVacation adds a vacation entry for children (presence module).
func PresenceAddVacation(ctx context.Context, s *aulaapi.Session, args *models.VacationEntry) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=calendar.addVacation", args)
}

// GetChildrenVacation gets children vacation overview (employee view).
func GetChildrenVacation(ctx context.Context, s *aulaapi.Session, args *models.ChildrenVacationRequest) (models.ChildrenVacationResult, error) {
	var query []string
	query = append(query, fmt.Sprintf("departmentId=%d", args.DepartmentID))
	if args.GroupIDs != nil {
		for _, id := range args.GroupIDs {
			query = append(query, fmt.Sprintf("groupIds=%d", id))
		}
	}
	if args.Date != nil {
		query = append(query, fmt.Sprintf("date=%s", EncodeValue(*args.Date)))
	}
	query = append(query, fmt.Sprintf("offset=%d", args.Offset))
	query = append(query, fmt.Sprintf("limit=%d", args.Limit))
	path := "?method=presence.getChildVacationList&" + strings.Join(query, "&")
	return aulaapi.SessionGet[models.ChildrenVacationResult](ctx, s, path)
}

// GetVacationAnnouncementsByChildren gets vacation announcements grouped by children.
func GetVacationAnnouncementsByChildren(ctx context.Context, s *aulaapi.Session, instProfileIDs []int64) ([]models.VacationAnnouncementsByChildren, error) {
	var query []string
	for _, id := range instProfileIDs {
		query = append(query, fmt.Sprintf("instProfileIds=%d", id))
	}
	path := "?method=presence.getVacationAnnouncementsByChildren"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.VacationAnnouncementsByChildren](ctx, s, path)
}

// GetVacationRegistrationOverview gets vacation registration overview (employee view).
func GetVacationRegistrationOverview(ctx context.Context, s *aulaapi.Session, args *models.ComeGoGetVacationRegistrationOverviewRequest) (models.GetVacationRegistrationOverview, error) {
	var query []string
	query = append(query, fmt.Sprintf("departmentId=%d", args.DepartmentID))
	if args.FilterGroups != nil {
		for _, g := range args.FilterGroups {
			query = append(query, fmt.Sprintf("filterGroups=%d", g))
		}
	}
	if args.StatusFilters != nil {
		for _, f := range args.StatusFilters {
			query = append(query, fmt.Sprintf("statusFilters=%s", EncodeValue(f)))
		}
	}
	query = append(query, fmt.Sprintf("offset=%d", args.Offset))
	query = append(query, fmt.Sprintf("limit=%d", args.Limit))
	path := "?method=presence.getVacationRegistrations&" + strings.Join(query, "&")
	return aulaapi.SessionGet[models.GetVacationRegistrationOverview](ctx, s, path)
}

// GetVacationRegistrationsByChildren gets vacation registrations grouped by children.
func GetVacationRegistrationsByChildren(ctx context.Context, s *aulaapi.Session, instProfileIDs []int64) ([]models.VacationRegistrationsByChildren, error) {
	var query []string
	for _, id := range instProfileIDs {
		query = append(query, fmt.Sprintf("instProfileIds=%d", id))
	}
	path := "?method=presence.getVacationRegistrationsByChildren"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.VacationRegistrationsByChildren](ctx, s, path)
}

// GetExistingVacationRegistrationResponse gets existing vacation registration response for a child.
func GetExistingVacationRegistrationResponse(ctx context.Context, s *aulaapi.Session, childID, vacationRegistrationID int64) (models.VacationRegistrationResponseForGuardian, error) {
	path := fmt.Sprintf("?method=presence.getVacationRegistrationResponse&childId=%d&vacationRegistrationId=%d", childID, vacationRegistrationID)
	return aulaapi.SessionGet[models.VacationRegistrationResponseForGuardian](ctx, s, path)
}

// GetPresenceConfiguration gets presence configuration.
func GetPresenceConfiguration(ctx context.Context, s *aulaapi.Session, institutionCode string) (models.PresenceConfigurationResult, error) {
	return aulaapi.SessionGet[models.PresenceConfigurationResult](ctx, s, fmt.Sprintf("?method=presence.getPresenceConfiguration&institutionCode=%s", institutionCode))
}

// GetPresenceConfigurationByChildrenIDs gets presence configuration by children IDs.
func GetPresenceConfigurationByChildrenIDs(ctx context.Context, s *aulaapi.Session, instProfileIDs []int64) ([]models.PresenceConfigurationChildResult, error) {
	var query []string
	for _, id := range instProfileIDs {
		query = append(query, fmt.Sprintf("instProfileIds=%d", id))
	}
	path := "?method=presence.getPresenceConfigurationByChildIds"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.PresenceConfigurationChildResult](ctx, s, path)
}

// GetPresenceFilter gets a single presence filter.
func GetPresenceFilter(ctx context.Context, s *aulaapi.Session, institutionCode string) (models.PresenceFilterResult, error) {
	return aulaapi.SessionGet[models.PresenceFilterResult](ctx, s, fmt.Sprintf("?method=presence.getPresenceFilters&institutionCode=%s", institutionCode))
}

// GetPresenceFilters gets multiple presence filters.
func GetPresenceFilters(ctx context.Context, s *aulaapi.Session, args *models.PresenceFiltersRequest) ([]models.PresenceFilterResult, error) {
	var query []string
	if args.Institutions != nil {
		for _, inst := range args.Institutions {
			query = append(query, fmt.Sprintf("institutions=%s", EncodeValue(inst)))
		}
	}
	path := "?method=presence.getPresenceFilters"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.PresenceFilterResult](ctx, s, path)
}

// GetClosedDays gets closed days.
func GetClosedDays(ctx context.Context, s *aulaapi.Session, institutionCodes []string) (models.GetClosedDaysResult, error) {
	var query []string
	for _, code := range institutionCodes {
		query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
	}
	path := "?method=presence.getClosedDays"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetClosedDaysResult](ctx, s, path)
}

// GetGeneralOpeningHours gets general opening hours.
func GetGeneralOpeningHours(ctx context.Context, s *aulaapi.Session, institutionCodes []string) (models.GetGeneralOpeningHoursResult, error) {
	var query []string
	for _, code := range institutionCodes {
		query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
	}
	path := "?method=presence.getGeneralOpeningHours"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetGeneralOpeningHoursResult](ctx, s, path)
}

// GetOpeningHoursByInstitutionCodes gets opening hours by institution codes within a date range.
func GetOpeningHoursByInstitutionCodes(ctx context.Context, s *aulaapi.Session, args *models.GetOpeningHoursByInstitutionCodesRequest) (models.GetOpeningHoursByInstitutionCodesResult, error) {
	var query []string
	if args.InstitutionCodes != nil {
		for _, code := range args.InstitutionCodes {
			query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
		}
	}
	if args.StartDate != nil {
		query = append(query, fmt.Sprintf("startDate=%s", EncodeValue(*args.StartDate)))
	}
	if args.EndDate != nil {
		query = append(query, fmt.Sprintf("endDate=%s", EncodeValue(*args.EndDate)))
	}
	path := "?method=presence.getOpeningHoursByInstitutionCodes"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetOpeningHoursByInstitutionCodesResult](ctx, s, path)
}

// GetSpecificOpeningHourOverview gets specific opening hour overview.
func GetSpecificOpeningHourOverview(ctx context.Context, s *aulaapi.Session, institutionCodes []string) (models.GetSpecificOpeningHourOverviewResult, error) {
	var query []string
	for _, code := range institutionCodes {
		query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
	}
	path := "?method=presence.getSpecificOpeningHourOverview"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[models.GetSpecificOpeningHourOverviewResult](ctx, s, path)
}

// GetAvailablePresenceStatuses gets available presence statuses.
func GetAvailablePresenceStatuses(ctx context.Context, s *aulaapi.Session, institutionCode string) (models.GetAvailableStatusesResult, error) {
	return aulaapi.SessionGet[models.GetAvailableStatusesResult](ctx, s, fmt.Sprintf("?method=presence.getPresenceStates&institutionCode=%s", institutionCode))
}

// GetInstitutionWithPresenceStates gets institutions with presence states.
func GetInstitutionWithPresenceStates(ctx context.Context, s *aulaapi.Session, institutionCodes []string) ([]models.InstitutionWithPresenceStates, error) {
	var query []string
	for _, code := range institutionCodes {
		query = append(query, fmt.Sprintf("institutionCodes=%s", EncodeValue(code)))
	}
	path := "?method=presence.getPresenceStates"
	if len(query) > 0 {
		path += "&" + strings.Join(query, "&")
	}
	return aulaapi.SessionGet[[]models.InstitutionWithPresenceStates](ctx, s, path)
}

// GetPresenceChildrenDistribution gets presence children distribution.
func GetPresenceChildrenDistribution(ctx context.Context, s *aulaapi.Session, args *models.PresenceChildrenDistributionRequestDto) (models.PresenceChildrenDistribution, error) {
	var query []string
	query = append(query, fmt.Sprintf("departmentId=%d", args.DepartmentID))
	if args.Date != nil {
		query = append(query, fmt.Sprintf("date=%s", EncodeValue(*args.Date)))
	}
	if args.GroupIDs != nil {
		for _, id := range args.GroupIDs {
			query = append(query, fmt.Sprintf("groupIds=%d", id))
		}
	}
	if args.StatusFilters != nil {
		for _, f := range args.StatusFilters {
			query = append(query, fmt.Sprintf("statusFilters=%s", EncodeValue(f)))
		}
	}
	path := "?method=presence.getPresenceDistribution&" + strings.Join(query, "&")
	return aulaapi.SessionGet[models.PresenceChildrenDistribution](ctx, s, path)
}
