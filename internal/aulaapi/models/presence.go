package models

import (
	"encoding/json"

	"github.com/lkt82/go-aula/internal/aulaapi/enums"
)

// PresenceRegistrationResult represents a presence registration result for a child.
type PresenceRegistrationResult struct {
	ID                     int64                       `json:"id"`
	InstitutionProfile     *SimpleInstitutionProfile   `json:"institutionProfile,omitempty"`
	Status                 *string                     `json:"status,omitempty"`
	ActivityType           *string                     `json:"activityType,omitempty"`
	Location               *ComeGoLocation             `json:"location,omitempty"`
	SleepIntervals         []SleepIntervalResult       `json:"sleepIntervals,omitempty"`
	EditablePresenceStates []string                    `json:"editablePresenceStates,omitempty"`
	CheckInTime            *string                     `json:"checkInTime,omitempty"`
	CheckOutTime           *string                     `json:"checkOutTime,omitempty"`
	SelfDeciderStartTime   *string                     `json:"selfDeciderStartTime,omitempty"`
	SelfDeciderEndTime     *string                     `json:"selfDeciderEndTime,omitempty"`
	EntryTime              *string                     `json:"entryTime,omitempty"`
	ExitTime               *string                     `json:"exitTime,omitempty"`
	ExitWith               *string                     `json:"exitWith,omitempty"`
	IsDefaultEntryTime     bool                        `json:"isDefaultEntryTime"`
	IsDefaultExitTime      bool                        `json:"isDefaultExitTime"`
	Comment                *string                     `json:"comment,omitempty"`
	SpareTimeActivity      *SpareTimeActivity          `json:"spareTimeActivity,omitempty"`
	VacationNote           *string                     `json:"vacationNote,omitempty"`
}

// ChildStatus represents a child status in the ComeGo system.
type ChildStatus struct {
	InstitutionProfileID int64                    `json:"institutionProfileId"`
	State                *string                  `json:"state,omitempty"`
	UniStudent           *ComeGoUniStudentProfile `json:"uniStudent,omitempty"`
}

// ChildStatusDto represents a child status DTO.
type ChildStatusDto struct {
	UniStudentID int64                    `json:"uniStudentId"`
	UniStudent   *ComeGoUniStudentProfile `json:"uniStudent,omitempty"`
	State        *enums.PresenceStatus    `json:"state,omitempty"`
}

// ComeGoUniStudentProfile represents a ComeGo-specific student profile.
type ComeGoUniStudentProfile struct {
	ID             *int64                         `json:"id,omitempty"`
	Name           *string                        `json:"name,omitempty"`
	ShortName      *string                        `json:"shortName,omitempty"`
	ProfilePicture *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
}

// ComeGoLocation represents a location result in ComeGo.
type ComeGoLocation struct {
	ID     int64   `json:"id"`
	Name   *string `json:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// PresenceLocation represents a presence location with description.
type PresenceLocation struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Symbol      *string `json:"symbol,omitempty"`
}

// PhysicalLocation represents a physical location in activity list.
type PhysicalLocation struct {
	ID   int64   `json:"id"`
	Name *string `json:"name,omitempty"`
}

// SleepIntervalResult represents a sleep interval result.
type SleepIntervalResult struct {
	ID        int64   `json:"id"`
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
}

// SpareTimeActivity represents spare time activity details.
type SpareTimeActivity struct {
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
	Comment   *string `json:"comment,omitempty"`
}

// InstitutionWithPresenceStates represents an institution with available presence states.
type InstitutionWithPresenceStates struct {
	InstitutionCode *string  `json:"institutionCode,omitempty"`
	PresenceStates  []string `json:"presenceStates,omitempty"`
}

// DateTimePeriod represents a date/time period.
type DateTimePeriod struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
}

// PresenceDay represents interface-like base for presence day data.
type PresenceDay struct {
	ID                   *int64             `json:"id,omitempty"`
	EntryTime            *string            `json:"entryTime,omitempty"`
	ExitTime             *string            `json:"exitTime,omitempty"`
	ExitWith             *string            `json:"exitWith,omitempty"`
	ByDate               *string            `json:"byDate,omitempty"`
	Comment              *string            `json:"comment,omitempty"`
	IsDefaultEntryTime   bool               `json:"isDefaultEntryTime"`
	IsDefaultExitTime    bool               `json:"isDefaultExitTime"`
	ActivityType         *string            `json:"activityType,omitempty"`
	SelfDeciderStartTime *string            `json:"selfDeciderStartTime,omitempty"`
	SelfDeciderEndTime   *string            `json:"selfDeciderEndTime,omitempty"`
	SpareTimeActivity    *SpareTimeActivity `json:"spareTimeActivity,omitempty"`
}

// PresenceDaySchedule represents a presence day schedule.
type PresenceDaySchedule struct {
	ID                                  *int64             `json:"id,omitempty"`
	EntryTime                           *string            `json:"entryTime,omitempty"`
	ExitTime                            *string            `json:"exitTime,omitempty"`
	ExitWith                            *string            `json:"exitWith,omitempty"`
	ByDate                              *string            `json:"byDate,omitempty"`
	Comment                             *string            `json:"comment,omitempty"`
	IsDefaultEntryTime                  bool               `json:"isDefaultEntryTime"`
	IsDefaultExitTime                   bool               `json:"isDefaultExitTime"`
	ActivityType                        *string            `json:"activityType,omitempty"`
	SelfDeciderStartTime                *string            `json:"selfDeciderStartTime,omitempty"`
	SelfDeciderEndTime                  *string            `json:"selfDeciderEndTime,omitempty"`
	SpareTimeActivity                   *SpareTimeActivity `json:"spareTimeActivity,omitempty"`
	DayOfWeek                           *string            `json:"dayOfWeek,omitempty"`
	FullName                            *string            `json:"fullName,omitempty"`
	DayText                             *string            `json:"dayText,omitempty"`
	RepeatPattern                       *string            `json:"repeatPattern,omitempty"`
	RepeatFromDate                      *string            `json:"repeatFromDate,omitempty"`
	RepeatToDate                        *string            `json:"repeatToDate,omitempty"`
	IsOnVacation                        bool               `json:"isOnVacation"`
	IsPlannedTimesOutsideOpeningHours   bool               `json:"isPlannedTimesOutsideOpeningHours"`
}

// GetDayTemplateResult represents a day template result.
type GetDayTemplateResult struct {
	CurrentDate            *string           `json:"currentDate,omitempty"`
	PresenceWeekTemplates  []json.RawMessage `json:"presenceWeekTemplates,omitempty"`
}

// GetAvailableStatusesResult represents available presence statuses result.
type GetAvailableStatusesResult struct {
	AvailableStatus []json.RawMessage `json:"availableStatus,omitempty"`
}

// ActivityListResult represents an activity list result with child counts.
type ActivityListResult struct {
	TotalNumberOfChildren   int                         `json:"totalNumberOfChildren"`
	NumberOfChildrenPresent int                         `json:"numberOfChildrenPresent"`
	Activities              []ActivityListChildPresence `json:"activities,omitempty"`
}

// ActivityListChildPresence represents a single child's presence in the activity list.
type ActivityListChildPresence struct {
	PresenceRegistrationID int64                    `json:"presenceRegistrationId"`
	UniStudent             *ActivityListChild       `json:"uniStudent,omitempty"`
	PresenceState          *string                  `json:"presenceState,omitempty"`
	Comment                *string                  `json:"comment,omitempty"`
	Note                   *string                  `json:"note,omitempty"`
	Location               *PresenceLocation        `json:"location,omitempty"`
	EditablePresenceStates []string                 `json:"editablePresenceStates,omitempty"`
	PastActivities         []PastPresenceActivity   `json:"pastActivities,omitempty"`
	FutureActivities       []FuturePresenceActivity `json:"futureActivities,omitempty"`
	IsEmphasized           bool                     `json:"isEmphasized"`
	IsDefaultEntryTimes    bool                     `json:"isDefaultEntryTimes"`
	IsDefaultExitTimes     bool                     `json:"isDefaultExitTimes"`
}

// ActivityListChild represents a child profile within the activity list.
type ActivityListChild struct {
	ID              int64                          `json:"id"`
	InstitutionCode *string                        `json:"institutionCode,omitempty"`
	ProfileID       int64                          `json:"profileId"`
	Role            *string                        `json:"role,omitempty"`
	ShortName       *string                        `json:"shortName,omitempty"`
	Metadata        *string                        `json:"metadata,omitempty"`
	Name            *string                        `json:"name,omitempty"`
	MainGroup       *string                        `json:"mainGroup,omitempty"`
	ProfilePicture  *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
}

// PastPresenceActivity represents a past presence activity record.
type PastPresenceActivity struct {
	CheckInTime  *string `json:"checkInTime,omitempty"`
	CheckoutTime *string `json:"checkoutTime,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	EndTime      *string `json:"endTime,omitempty"`
	ActivityType *string `json:"activityType,omitempty"`
}

// FuturePresenceActivity represents a future presence activity record.
type FuturePresenceActivity struct {
	ActivityType         *string `json:"activityType,omitempty"`
	EntryTime            *string `json:"entryTime,omitempty"`
	ExitTime             *string `json:"exitTime,omitempty"`
	ExitWith             *string `json:"exitWith,omitempty"`
	SelfDeciderStartTime *string `json:"selfDeciderStartTime,omitempty"`
	SelfDeciderEndTime   *string `json:"selfDeciderEndTime,omitempty"`
	StartTime            *string `json:"startTime,omitempty"`
	EndTime              *string `json:"endTime,omitempty"`
}

// ActivityListRequest represents activity list request parameters.
type ActivityListRequest struct {
	DepartmentID int64    `json:"departmentId"`
	GroupIDs     []int64  `json:"groupIds,omitempty"`
	Limit        *int     `json:"limit,omitempty"`
	Offset       *int     `json:"offset,omitempty"`
	States       []string `json:"states,omitempty"`
	NextActivity *string  `json:"nextActivity,omitempty"`
	LocationIDs  []int64  `json:"locationIds,omitempty"`
	SortOn       *string  `json:"sortOn,omitempty"`
	DailyNote    *string  `json:"dailyNote,omitempty"`
}

// ActivityFilterResult represents an activity filter.
type ActivityFilterResult struct {
	InstitutionCode        *string                    `json:"institutionCode,omitempty"`
	InstitutionName        *string                    `json:"institutionName,omitempty"`
	Departments            []PresenceFilterDepartment `json:"departments,omitempty"`
	PresenceStates         []string                   `json:"presenceStates,omitempty"`
	PresenceNextActivities []string                   `json:"presenceNextActivities,omitempty"`
	Locations              []PhysicalLocation         `json:"locations,omitempty"`
}

// InstitutionWithPresenceStatesResult represents an institution with presence states (view model variant).
type InstitutionWithPresenceStatesResult struct {
	InstitutionCode        *string  `json:"institutionCode,omitempty"`
	EditablePresenceStates []string `json:"editablePresenceStates,omitempty"`
}

// PresenceFilterResult represents a presence filter result (per institution).
type PresenceFilterResult struct {
	InstitutionCode *string                    `json:"institutionCode,omitempty"`
	InstitutionName *string                    `json:"institutionName,omitempty"`
	Departments     []PresenceFilterDepartment `json:"departments,omitempty"`
}

// PresenceFilterDepartment represents a filter department within a presence filter.
type PresenceFilterDepartment struct {
	ID              int64                `json:"id"`
	FilteringGroups []PresenceFilterGroup `json:"filteringGroups,omitempty"`
	MainGroup       *MainGroup           `json:"mainGroup,omitempty"`
	Name            *string              `json:"name,omitempty"`
	IsSelected      bool                 `json:"isSelected"`
}

// PresenceFilterGroup represents a filter group within a department.
type PresenceFilterGroup struct {
	ID         int64   `json:"id"`
	Name       *string `json:"name,omitempty"`
	IsSelected bool    `json:"isSelected"`
}

// PresenceFiltersRequest represents a presence filters request.
type PresenceFiltersRequest struct {
	Institutions []string `json:"institutions,omitempty"`
}

// PresenceSchedulesRequest represents a request for presence schedules.
type PresenceSchedulesRequest struct {
	FilterInstitutionProfileIDs []int64 `json:"filterInstitutionProfileIds,omitempty"`
	FromDate                    *string `json:"fromDate,omitempty"`
	ToDate                      *string `json:"toDate,omitempty"`
}

// GetPickupResponsibleChildResult represents a result for a child's pickup responsible persons.
type GetPickupResponsibleChildResult struct {
	UniStudentID      int64                                `json:"uniStudentId"`
	RelatedPersons    []PresenceRelatedPersonPickResponsible `json:"relatedPersons,omitempty"`
	PickupSuggestions []PresencePickupSuggestion           `json:"pickupSuggestions,omitempty"`
}

// GetPickupResponsibleResult represents the top-level pickup responsible result.
type GetPickupResponsibleResult struct {
	Children []GetPickupResponsibleChildResult `json:"children,omitempty"`
}

// PresenceRelatedPersonPickResponsible represents a related person who can pick up a child.
type PresenceRelatedPersonPickResponsible struct {
	InstitutionProfileID *int64  `json:"institutionProfileId,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Relation             *string `json:"relation,omitempty"`
}

// PresencePickupSuggestion represents a pickup name suggestion.
type PresencePickupSuggestion struct {
	ID           int64   `json:"id"`
	UniStudentID int64   `json:"uniStudentId"`
	PickUpName   *string `json:"pickUpName,omitempty"`
}

// GetPickupResponsibleRequest represents a request to get pickup responsible.
type GetPickupResponsibleRequest struct {
	UniStudentIDs []int64 `json:"uniStudentIds,omitempty"`
}

// SavePickupNameRequest represents a request to save a pickup name.
type SavePickupNameRequest struct {
	ID   int64   `json:"id"`
	Name *string `json:"name,omitempty"`
}

// DeletePickupResponsibleRequest represents a request to delete a pickup responsible entry.
type DeletePickupResponsibleRequest struct {
	PresencePickupSuggestionID int64 `json:"presencePickupSuggestionId"`
}

// ComeGoExitWithSuggestion represents an exit-with suggestion.
type ComeGoExitWithSuggestion struct {
	PickupName   *string `json:"pickupName,omitempty"`
	UniStudentID int64   `json:"uniStudentId"`
}

// ComeGoExitWithSuggestionRequest represents a request for exit-with suggestions.
type ComeGoExitWithSuggestionRequest struct {
	PickupName    *string `json:"pickupName,omitempty"`
	UniStudentIDs []int64 `json:"uniStudentIds,omitempty"`
}

// GetExitWithSuggestionsResult represents a result wrapper for exit-with suggestions.
type GetExitWithSuggestionsResult struct {
	Suggestions []ComeGoExitWithSuggestion `json:"suggestions,omitempty"`
}

// AddSleepIntervalsRequest represents a request to add sleep intervals.
type AddSleepIntervalsRequest struct {
	ChildIDs []int64 `json:"childIds,omitempty"`
	Start    *string `json:"start,omitempty"`
	End      *string `json:"end,omitempty"`
}

// UpdateSleepIntervals represents a request to update sleep intervals.
type UpdateSleepIntervals struct {
	SleepIntervalIDs []int64 `json:"sleepIntervalIds,omitempty"`
	Start            *string `json:"start,omitempty"`
	End              *string `json:"end,omitempty"`
}

// UpdateSleepIntervalsDto represents a DTO for updating sleep intervals.
type UpdateSleepIntervalsDto struct {
	PresenceRegistrationID int64   `json:"presenceRegistrationId"`
	ID                     int64   `json:"id"`
	Start                  *string `json:"start,omitempty"`
	End                    *string `json:"end,omitempty"`
}

// OpeningHours represents general opening hours for a day of the week.
type OpeningHours struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	DayOfWeek       *string `json:"dayOfWeek,omitempty"`
	OpenTime        *string `json:"openTime,omitempty"`
	CloseTime       *string `json:"closeTime,omitempty"`
}

// InstitutionOpeningHours represents institution opening hours wrapper.
type InstitutionOpeningHours struct {
	InstitutionCode *string        `json:"institutionCode,omitempty"`
	OpeningHours    []OpeningHours `json:"openingHours,omitempty"`
}

// GetGeneralOpeningHoursResult represents general opening hours result.
type GetGeneralOpeningHoursResult struct {
	InstitutionOpeningHours []InstitutionOpeningHours `json:"institutionOpeningHours,omitempty"`
}

// OpeningHoursDto represents opening hours for a specific date/period.
type OpeningHoursDto struct {
	InstitutionCode  *string `json:"institutionCode,omitempty"`
	Date             *string `json:"date,omitempty"`
	OpenTime         *string `json:"openTime,omitempty"`
	CloseTime        *string `json:"closeTime,omitempty"`
	Name             *string `json:"name,omitempty"`
	OpeningHoursType *string `json:"type,omitempty"`
}

// OpeningHoursOverview represents opening hours overview per institution.
type OpeningHoursOverview struct {
	InstitutionCode *string           `json:"institutionCode,omitempty"`
	OpeningHoursDto []OpeningHoursDto `json:"openingHoursDto,omitempty"`
}

// GetOpeningHoursByInstitutionCodesRequest represents a request for opening hours by institution codes.
type GetOpeningHoursByInstitutionCodesRequest struct {
	InstitutionCodes []string `json:"institutionCodes,omitempty"`
	StartDate        *string  `json:"startDate,omitempty"`
	EndDate          *string  `json:"endDate,omitempty"`
}

// GetOpeningHoursByInstitutionCodesResult represents a result for opening hours by institution codes.
type GetOpeningHoursByInstitutionCodesResult struct {
	OpeningHoursOverviewDto []OpeningHoursOverview `json:"openingHoursOverviewDto,omitempty"`
}

// SpecificOpeningHour represents a specific opening hour entry.
type SpecificOpeningHour struct {
	ID        int64   `json:"id"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
	OpenTime  *string `json:"openTime,omitempty"`
	CloseTime *string `json:"closeTime,omitempty"`
}

// SpecificOpeningHourWithInstitution represents specific opening hours per institution.
type SpecificOpeningHourWithInstitution struct {
	InstitutionCode      *string               `json:"institutionCode,omitempty"`
	SpecificOpeningHours []SpecificOpeningHour `json:"specificOpeningHours,omitempty"`
}

// GetSpecificOpeningHourOverviewResult represents result for specific opening hours overview.
type GetSpecificOpeningHourOverviewResult struct {
	SpecificOpeningHoursWithInstitutions []SpecificOpeningHourWithInstitution `json:"specificOpeningHoursWithInstitutions,omitempty"`
}

// ClosedDay represents a closed day entry.
type ClosedDay struct {
	ID        int64   `json:"id"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// ClosedDaysOverview represents a closed days overview.
type ClosedDaysOverview struct {
	ClosedDays []ClosedDay `json:"closedDays,omitempty"`
}

// InstitutionClosedDays represents institution closed days.
type InstitutionClosedDays struct {
	InstitutionCode    *string             `json:"institutionCode,omitempty"`
	ClosedDaysOverview *ClosedDaysOverview `json:"closedDaysOverview,omitempty"`
}

// GetClosedDaysResult represents result for closed days query.
type GetClosedDaysResult struct {
	InstitutionClosedDays []InstitutionClosedDays `json:"institutionClosedDays,omitempty"`
}

// PresenceConfigurationChildResult represents configuration result for a child's presence.
type PresenceConfigurationChildResult struct {
	UniStudentID           int64                         `json:"uniStudentId"`
	PresenceConfiguration  *PresenceConfigurationResult  `json:"presenceConfiguration,omitempty"`
}

// PresenceConfigurationResult represents full presence configuration result.
type PresenceConfigurationResult struct {
	SelfDecider              bool                              `json:"selfDecider"`
	GoHomeWith               bool                              `json:"goHomeWith"`
	SendHome                 bool                              `json:"sendHome"`
	PickUp                   bool                              `json:"pickUp"`
	Institution              *PresenceConfigurationInstitution `json:"institution,omitempty"`
	Departments              []PresenceConfigurationDepartment `json:"departments,omitempty"`
	DashboardModuleSettings  []PresenceModuleSettings          `json:"dashboardModuleSettings,omitempty"`
}

// PresenceConfigurationInstitution represents institution identity within presence configuration.
type PresenceConfigurationInstitution struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	Name            *string `json:"name,omitempty"`
}

// PresenceConfigurationDepartment represents department within presence configuration.
type PresenceConfigurationDepartment struct {
	Group           *PresenceConfigurationGroup   `json:"group,omitempty"`
	FilteringGroups []PresenceConfigurationGroup  `json:"filteringGroups,omitempty"`
}

// PresenceConfigurationGroup represents group within presence configuration.
type PresenceConfigurationGroup struct {
	ID   int    `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PresenceModule represents a presence module with type and permission.
type PresenceModule struct {
	ModuleType *string `json:"moduleType,omitempty"`
	Permission *string `json:"permission,omitempty"`
}

// PresenceModuleSettings represents presence module settings per dashboard context.
type PresenceModuleSettings struct {
	PresenceDashboardContext *string          `json:"presenceDashboardContext,omitempty"`
	PresenceModules         []PresenceModule `json:"presenceModules,omitempty"`
}

// ParentDailyOverviewInstitutionProfile represents parent daily overview institution profile.
type ParentDailyOverviewInstitutionProfile struct {
	ProfileID       int64                          `json:"profileId"`
	ID              int64                          `json:"id"`
	InstitutionCode *string                        `json:"institutionCode,omitempty"`
	InstitutionName *string                        `json:"institutionName,omitempty"`
	Role            *string                        `json:"role,omitempty"`
	Name            *string                        `json:"name,omitempty"`
	ProfilePicture  *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
	MainGroup       *MainGroup                     `json:"mainGroup,omitempty"`
	ShortName       *string                        `json:"shortName,omitempty"`
	InstitutionRole *string                        `json:"institutionRole,omitempty"`
	Metadata        *string                        `json:"metadata,omitempty"`
}

// ParentsDailyOverviewResult represents parent daily overview result for a child.
type ParentsDailyOverviewResult struct {
	InstitutionProfile                  *ParentDailyOverviewInstitutionProfile `json:"institutionProfile,omitempty"`
	MainGroup                           *MainGroup                            `json:"mainGroup,omitempty"`
	Status                              *string                               `json:"status,omitempty"`
	SleepIntervals                      []SleepIntervalResult                 `json:"sleepIntervals,omitempty"`
	CheckInTime                         *string                               `json:"checkInTime,omitempty"`
	CheckOutTime                        *string                               `json:"checkOutTime,omitempty"`
	Location                            *PresenceLocation                     `json:"location,omitempty"`
	IsDefaultEntryTime                  bool                                  `json:"isDefaultEntryTime"`
	IsDefaultExitTime                   bool                                  `json:"isDefaultExitTime"`
	IsPlannedTimesOutsideOpeningHours   bool                                  `json:"isPlannedTimesOutsideOpeningHours"`
}

// VacationAnnouncement represents a vacation announcement.
type VacationAnnouncement struct {
	VacationID  int64   `json:"vacationId"`
	StartDate   *string `json:"startDate,omitempty"`
	EndDate     *string `json:"endDate,omitempty"`
	Description *string `json:"description,omitempty"`
	IsEditable  bool    `json:"isEditable"`
}

// VacationAnnouncementsByChildren represents vacation announcements grouped by child.
type VacationAnnouncementsByChildren struct {
	Child                  *ParentDailyOverviewInstitutionProfile `json:"child,omitempty"`
	VacationAnnouncements  []VacationAnnouncement                `json:"vacationAnnouncements,omitempty"`
}

// PresenceVacationRegistration represents a vacation registration (staff-created) in presence context.
type PresenceVacationRegistration struct {
	VacationRegistrationID int64   `json:"vacationRegistrationId"`
	StartDate              *string `json:"startDate,omitempty"`
	EndDate                *string `json:"endDate,omitempty"`
	Title                  *string `json:"title,omitempty"`
	NoteToGuardian         *string `json:"noteToGuardian,omitempty"`
	ResponseID             int64   `json:"responseId"`
	ResponseDeadline       *string `json:"responseDeadline,omitempty"`
	IsEditable             bool    `json:"isEditable"`
	IsMissingAnswer        bool    `json:"isMissingAnswer"`
	IsPresenceTimesRequired bool   `json:"isPresenceTimesRequired"`
}

// VacationRegistrationsByChildren represents vacation registrations grouped by child.
type VacationRegistrationsByChildren struct {
	Child                  *ParentDailyOverviewInstitutionProfile `json:"child,omitempty"`
	VacationRegistrations  []PresenceVacationRegistration         `json:"vacationRegistrations,omitempty"`
}

// VacationEntry represents a vacation entry (guardian submits intervals).
type VacationEntry struct {
	ChildIDs  []int64                        `json:"childIds,omitempty"`
	Intervals []PresenceVacationIntervals    `json:"intervals,omitempty"`
	Comment   *string                        `json:"comment,omitempty"`
}

// PresenceVacationIntervals represents vacation interval (date range).
type PresenceVacationIntervals struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
}

// PresenceGuardianRegisterVacationIntervals represents guardian vacation registration intervals.
type PresenceGuardianRegisterVacationIntervals struct {
	Date      *string `json:"date,omitempty"`
	EntryTime *string `json:"entryTime,omitempty"`
	ExitTime  *string `json:"exitTime,omitempty"`
	IsComing  bool    `json:"isComing"`
}

// ChildrenVacationRequest represents children vacation request parameters.
type ChildrenVacationRequest struct {
	DepartmentID int64   `json:"departmentId"`
	GroupIDs     []int64 `json:"groupIds,omitempty"`
	Date         *string `json:"date,omitempty"`
	Offset       int     `json:"offset"`
	Limit        int     `json:"limit"`
}

// ChildrenVacationResult represents children vacation result (paginated).
type ChildrenVacationResult struct {
	Count    int                      `json:"count"`
	Children []ChildrenVacationChild `json:"children,omitempty"`
}

// ChildrenVacationChild represents a child within vacation result.
type ChildrenVacationChild struct {
	Child *ChildrenVacationChildProfile `json:"child,omitempty"`
	Note  *string                       `json:"note,omitempty"`
}

// ChildrenVacationChildProfile represents a child profile within vacation context.
type ChildrenVacationChildProfile struct {
	ProfileID      int64                          `json:"profileId"`
	ID             int64                          `json:"id"`
	ShortName      *string                        `json:"shortName,omitempty"`
	Name           *string                        `json:"name,omitempty"`
	Metadata       *string                        `json:"metadata,omitempty"`
	ProfilePicture *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
}

// VacationFilterResult represents a vacation filter.
type VacationFilterResult struct {
	InstitutionCode *string                    `json:"institutionCode,omitempty"`
	InstitutionName *string                    `json:"institutionName,omitempty"`
	Departments     []PresenceFilterDepartment `json:"departments,omitempty"`
}

// RespondToVacationRegistrationRequest represents responding to vacation registration request.
type RespondToVacationRegistrationRequest struct {
	ChildID                        *int64                                      `json:"childId,omitempty"`
	VacationRegistrationResponseID *int64                                      `json:"vacationRegistrationResponseId,omitempty"`
	Days                           []PresenceGuardianRegisterVacationIntervals `json:"days,omitempty"`
	Comment                        *string                                     `json:"comment,omitempty"`
}

// VacationRegistrationResponseForGuardian represents vacation registration response for guardian.
type VacationRegistrationResponseForGuardian struct {
	VacationRegistration         *PresenceVacationRegistration          `json:"vacationRegistration,omitempty"`
	VacationRegistrationResponse *RespondToVacationRegistrationRequest  `json:"vacationRegistrationResponse,omitempty"`
}

// CreateVacationRegistrationRequest represents creating a vacation registration request (staff).
type CreateVacationRegistrationRequest struct {
	StartDateTime        *string                        `json:"startDateTime,omitempty"`
	EndDateTime          *string                        `json:"endDateTime,omitempty"`
	ResponseDeadline     *string                        `json:"responseDeadline,omitempty"`
	CreatorInstProfileID int64                          `json:"creatorInstProfileId"`
	Title                *string                        `json:"title,omitempty"`
	Departments          []DepartmentIdsSimpleRequest   `json:"departments,omitempty"`
	NoteToGuardians      *string                        `json:"noteToGuardians,omitempty"`
	IsPresenceTimesRequired bool                        `json:"isPresenceTimesRequired"`
}

// DepartmentIdsSimpleRequest represents a simple department ID reference.
type DepartmentIdsSimpleRequest struct {
	GroupID         int64   `json:"groupId"`
	FilteringGroups []int64 `json:"filteringGroups,omitempty"`
}

// UpdateVacationRegistrationRequest represents updating a vacation registration request.
type UpdateVacationRegistrationRequest struct {
	ID               int64   `json:"id"`
	ResponseDeadline *string `json:"responseDeadline,omitempty"`
}

// ComeGoPresenceTimeWithTense represents a presence time with tense.
type ComeGoPresenceTimeWithTense struct {
	Timestamp *string `json:"timestamp,omitempty"`
	Tense     *string `json:"tense,omitempty"`
}

// EmployeeWeekOverviewPresenceDetails represents presence details for employee week overview.
type EmployeeWeekOverviewPresenceDetails struct {
	StartTime *ComeGoPresenceTimeWithTense `json:"startTime,omitempty"`
	EndTime   *ComeGoPresenceTimeWithTense `json:"endTime,omitempty"`
}

// EmployeeWeekOverviewVacationDetails represents vacation details for employee week overview.
type EmployeeWeekOverviewVacationDetails struct {
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
}

// EmployeeWeekOverviewActivities represents activities for a single day in week overview.
type EmployeeWeekOverviewActivities struct {
	Date            *string                              `json:"date,omitempty"`
	PresenceType    *string                              `json:"type,omitempty"`
	PresenceDetails *EmployeeWeekOverviewPresenceDetails `json:"presenceDetails,omitempty"`
	VacationDetails *EmployeeWeekOverviewVacationDetails `json:"vacationDetails,omitempty"`
}

// EmployeeWeekOverviewChildActivities represents child activities in employee week overview.
type EmployeeWeekOverviewChildActivities struct {
	Child                                       *ActivityListChild                `json:"child,omitempty"`
	Activities                                  []EmployeeWeekOverviewActivities  `json:"activities,omitempty"`
	PresenceRegistrationID                      *int64                            `json:"presenceRegistrationId,omitempty"`
	PresenceRegistrationIsDefaultEntryTime       *bool                             `json:"presenceRegistrationIsDefaultEntryTime,omitempty"`
	PresenceRegistrationIsDefaultExitTime        *bool                             `json:"presenceRegistrationIsDefaultExitTime,omitempty"`
}

// GetPresenceOverview represents the presence overview result for employee week view.
type GetPresenceOverview struct {
	WeekNumber      int                                   `json:"weekNumber"`
	PresenceDays    []WeekOverviewPresenceDays            `json:"presenceDays,omitempty"`
	ChildActivities []EmployeeWeekOverviewChildActivities `json:"childActivities,omitempty"`
}

// WeekOverviewPresenceDays represents a day summary in week overview.
type WeekOverviewPresenceDays struct {
	Date                   *string `json:"date,omitempty"`
	NumberOfChildren       int     `json:"numberOfChildren"`
	TotalNumberOfChildren  int     `json:"totalNumberOfChildren"`
}

// PresenceChildrenDistribution represents children distribution in a presence overview.
type PresenceChildrenDistribution struct {
	NumberPresent         int                     `json:"numberPresent"`
	NumberOnVacation      int                     `json:"numberOnVacation"`
	NumberSick            int                     `json:"numberSick"`
	NumberNotArrived      int                     `json:"numberNotArrived"`
	Intervals             []PresenceIntervalModel `json:"intervals,omitempty"`
	IsDistributionEnabled bool                    `json:"isDistributionEnabled"`
}

// PresenceIntervalModel represents a presence interval.
type PresenceIntervalModel struct {
	StartTime        *string `json:"startTime,omitempty"`
	EndTime          *string `json:"endTime,omitempty"`
	NumberOfChildren *string `json:"numberOfChildren,omitempty"`
	IsCurrent        bool    `json:"isCurrent"`
}

// ChildrenPresenceDistributionRequest represents children presence distribution request.
type ChildrenPresenceDistributionRequest struct {
	ShowingDate *string                                  `json:"showingDate,omitempty"`
	Dto         *PresenceChildrenDistributionRequestDto `json:"dto,omitempty"`
}

// PresenceChildrenDistributionRequestDto represents presence children distribution request DTO.
type PresenceChildrenDistributionRequestDto struct {
	DepartmentID  int64    `json:"departmentId"`
	Date          *string  `json:"date,omitempty"`
	GroupIDs      []int64  `json:"groupIds,omitempty"`
	StatusFilters []string `json:"statusFilters,omitempty"`
}

// ComeGoGetWeekOverviewRequest represents a request to get employee week overview.
type ComeGoGetWeekOverviewRequest struct {
	DepartmentID  int64    `json:"departmentId"`
	GroupIDs      []int64  `json:"groupIds,omitempty"`
	StatusFilters []string `json:"statusFilters,omitempty"`
	StartDate     *string  `json:"startDate,omitempty"`
	EndDate       *string  `json:"endDate,omitempty"`
	Offset        int      `json:"offset"`
	Limit         int      `json:"limit"`
}

// ComeGoGetVacationRegistrationOverviewRequest represents a request to get vacation registration overview.
type ComeGoGetVacationRegistrationOverviewRequest struct {
	DepartmentID  int64    `json:"departmentId"`
	FilterGroups  []int64  `json:"filterGroups,omitempty"`
	StatusFilters []string `json:"statusFilters,omitempty"`
	Offset        int      `json:"offset"`
	Limit         int      `json:"limit"`
}

// GetVacationRegistrationOverview represents vacation registration overview result.
type GetVacationRegistrationOverview struct {
	TotalNumber            int                      `json:"totalNumber"`
	VacationRegistrations  []VacationRegistrationsDto `json:"vacationRegistrations,omitempty"`
}

// VacationRegistrationsDto represents a vacation registration entry in overview list.
type VacationRegistrationsDto struct {
	VacationRegistrationID                int    `json:"vacationRegistrationId"`
	Title                                 *string `json:"title,omitempty"`
	StartDate                             *string `json:"startDate,omitempty"`
	EndDate                               *string `json:"endDate,omitempty"`
	Deadline                              *string `json:"deadline,omitempty"`
	RegardingDepartmentAndGroupsText      []string `json:"regardingDepartmentAndGroupsText,omitempty"`
	Subtitle                              *string `json:"subtitle,omitempty"`
	ShortName                             *string `json:"shortName,omitempty"`
}

// OverallItem represents an overall item in employee week overview.
type OverallItem struct {
	LeftText              *string `json:"leftText,omitempty"`
	LeftTextAccessibility *string `json:"leftTextAccessibility,omitempty"`
	RightText             *string `json:"rightText,omitempty"`
}

// EmployeeWeekOverviewPresence represents employee week overview presence record.
type EmployeeWeekOverviewPresence struct {
	ActivityType         *string            `json:"activityType,omitempty"`
	ByDate               *string            `json:"byDate,omitempty"`
	Comment              *string            `json:"comment,omitempty"`
	DayOfWeek            *string            `json:"dayOfWeek,omitempty"`
	EntryTime            *string            `json:"entryTime,omitempty"`
	ExitTime             *string            `json:"exitTime,omitempty"`
	ExitWith             *string            `json:"exitWith,omitempty"`
	ID                   *int64             `json:"id,omitempty"`
	IsOnVacation         bool               `json:"isOnVacation"`
	IsRepeating          bool               `json:"isRepeating"`
	RepeatFromDate       *string            `json:"repeatFromDate,omitempty"`
	RepeatToDate         *string            `json:"repeatToDate,omitempty"`
	SelfDeciderEndTime   *string            `json:"selfDeciderEndTime,omitempty"`
	SelfDeciderStartTime *string            `json:"selfDeciderStartTime,omitempty"`
	SpareTimeActivity    *SpareTimeActivity `json:"spareTimeActivity,omitempty"`
	Vacation             json.RawMessage    `json:"vacation,omitempty"`
}

// WeekOverviewFutureDate represents a week overview future date model.
type WeekOverviewFutureDate struct {
	EntryTime            *string `json:"entryTime,omitempty"`
	ExitTime             *string `json:"exitTime,omitempty"`
	SelfDeciderStartTime *string `json:"selfDeciderStartTime,omitempty"`
	SelfDeciderEndTime   *string `json:"selfDeciderEndTime,omitempty"`
	ExitWith             *string `json:"exitWith,omitempty"`
	ActivityType         *string `json:"activityType,omitempty"`
}

// PresenceRegistrationRequest represents a presence registration request.
type PresenceRegistrationRequest struct {
	ChildID int64   `json:"childId"`
	Date    *string `json:"date,omitempty"`
}

// PresenceRegistrationTodayRequest represents a presence registration today request.
type PresenceRegistrationTodayRequest struct {
	PresenceRegistrationIDs []int64 `json:"presenceRegistrationIds,omitempty"`
	DepartmentID            *string `json:"departmentId,omitempty"`
}

// BulkUpdatePresenceStatusRequest represents a bulk update presence status request.
type BulkUpdatePresenceStatusRequest struct {
	PresenceRegistrationIDs []int64 `json:"presenceRegistrationIds,omitempty"`
	Status                  *string `json:"status,omitempty"`
}

// UpdateStatusByInstitutionProfileIds represents updating status by institution profile IDs.
type UpdateStatusByInstitutionProfileIds struct {
	InstitutionProfileIDs []int64 `json:"institutionProfileIds,omitempty"`
	Status                int     `json:"status"`
}

// UpdateStatus represents updating status by IDs.
type UpdateStatus struct {
	IDs    []int64 `json:"ids,omitempty"`
	Status int     `json:"status"`
}

// UpdatePresenceRegistrationRequest represents an update presence registration request.
type UpdatePresenceRegistrationRequest struct {
	RegistrationID int64                             `json:"registrationId"`
	CheckoutType   *string                           `json:"checkoutType,omitempty"`
	PickupBy       *UpdateCheckoutPickedUpActivity   `json:"pickupBy,omitempty"`
	SelfDecider    *UpdateCheckoutSelfDeciderActivity `json:"selfDecider,omitempty"`
	SendHome       *UpdateCheckoutSendHomeActivity   `json:"sendHome,omitempty"`
	GoHomeWith     *UpdateCheckoutGoHomeWithActivity  `json:"goHomeWith,omitempty"`
	EntryTime      *string                           `json:"entryTime,omitempty"`
	Remark         *string                           `json:"remark,omitempty"`
}

// UpdateCheckoutPickedUpActivity represents checkout picked-up activity.
type UpdateCheckoutPickedUpActivity struct {
	ExitTime *string `json:"exitTime,omitempty"`
	ExitWith *string `json:"exitWith,omitempty"`
}

// UpdateCheckoutSelfDeciderActivity represents checkout self-decider activity.
type UpdateCheckoutSelfDeciderActivity struct {
	SelfDeciderStartTime *string `json:"selfDeciderStartTime,omitempty"`
	SelfDeciderEndTime   *string `json:"selfDeciderEndTime,omitempty"`
}

// UpdateCheckoutSendHomeActivity represents checkout send-home activity.
type UpdateCheckoutSendHomeActivity struct {
	ExitTime *string `json:"exitTime,omitempty"`
}

// UpdateCheckoutGoHomeWithActivity represents checkout go-home-with activity.
type UpdateCheckoutGoHomeWithActivity struct {
	ExitWith *string `json:"exitWith,omitempty"`
	ExitTime *string `json:"exitTime,omitempty"`
}

// UpdatePresenceDayRequest represents an update presence day request.
type UpdatePresenceDayRequest struct {
	InstitutionProfileID int64                       `json:"institutionProfileId"`
	ID                   *int64                      `json:"id,omitempty"`
	DayOfWeek            int                         `json:"dayOfWeek"`
	ByDate               *string                     `json:"byDate,omitempty"`
	Comment              *string                     `json:"comment,omitempty"`
	SpareTimeActivity    *SpareTimeActivityRequest   `json:"spareTimeActivity,omitempty"`
	PresenceActivity     *UpdatePresenceDayActivity  `json:"presenceActivity,omitempty"`
	RepeatPattern        *string                     `json:"repeatPattern,omitempty"`
	ExpiresAt            *string                     `json:"expiresAt,omitempty"`
}

// SpareTimeActivityRequest represents spare time activity request model.
type SpareTimeActivityRequest struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
	Comment   *string `json:"comment,omitempty"`
}

// UpdatePresenceDayActivity represents an update presence day activity.
type UpdatePresenceDayActivity struct {
	ActivityType *string                                `json:"activityType,omitempty"`
	Pickup       *UpdatePresenceDayPickedUpActivity     `json:"pickup,omitempty"`
	SelfDecider  *UpdatePresenceDaySelfDeciderActivity  `json:"selfDecider,omitempty"`
	SendHome     *UpdatePresenceDaySendHomeActivity     `json:"sendHome,omitempty"`
	GoHomeWith   *UpdatePresenceDayGoHomeWithActivity   `json:"goHomeWith,omitempty"`
	EntryTime    *string                                `json:"entryTime,omitempty"`
	ExitTime     *string                                `json:"exitTime,omitempty"`
}

// UpdatePresenceDayPickedUpActivity represents picked-up activity in presence day update.
type UpdatePresenceDayPickedUpActivity struct {
	EntryTime *string `json:"entryTime,omitempty"`
	ExitTime  *string `json:"exitTime,omitempty"`
	ExitWith  *string `json:"exitWith,omitempty"`
}

// UpdatePresenceDaySelfDeciderActivity represents self-decider activity in presence day update.
type UpdatePresenceDaySelfDeciderActivity struct {
	EntryTime     *string `json:"entryTime,omitempty"`
	ExitStartTime *string `json:"exitStartTime,omitempty"`
	ExitEndTime   *string `json:"exitEndTime,omitempty"`
}

// UpdatePresenceDaySendHomeActivity represents send-home activity in presence day update.
type UpdatePresenceDaySendHomeActivity struct {
	EntryTime *string `json:"entryTime,omitempty"`
	ExitTime  *string `json:"exitTime,omitempty"`
}

// UpdatePresenceDayGoHomeWithActivity represents go-home-with activity in presence day update.
type UpdatePresenceDayGoHomeWithActivity struct {
	ExitWith  *string `json:"exitWith,omitempty"`
	EntryTime *string `json:"entryTime,omitempty"`
	ExitTime  *string `json:"exitTime,omitempty"`
}

// DeletePresenceTemplateRequest represents a delete presence template request.
type DeletePresenceTemplateRequest struct {
	DeleteFromDay      *string `json:"deleteFromDay,omitempty"`
	PresentTemplateID  *int64  `json:"presentTemplateId,omitempty"`
}

// GetOverlappingPresenceTemplatesRequest represents a get overlapping presence templates request.
type GetOverlappingPresenceTemplatesRequest struct {
	InstitutionProfileID int64   `json:"institutionProfileId"`
	StartDate            *string `json:"startDate,omitempty"`
	EndDate              *string `json:"endDate,omitempty"`
	RepeatPattern        *string `json:"repeatPattern,omitempty"`
}

// UpdateLocationRequest represents an update location request.
type UpdateLocationRequest struct {
	ChildIDs   []int64 `json:"childIds,omitempty"`
	LocationID *int64  `json:"locationId,omitempty"`
}

// UpdatePickUpResponsibleResult represents update pickup responsible result.
type UpdatePickUpResponsibleResult struct {
	Result             bool `json:"result"`
	HasWhiteSpaceError bool `json:"hasWhiteSpaceError"`
}

// ChildGoHomeWithResult represents child go-home-with result.
type ChildGoHomeWithResult struct {
	InstitutionProfileID int     `json:"institutionProfileId"`
	FullName             *string `json:"fullName,omitempty"`
	MainGroup            *string `json:"mainGroup,omitempty"`
}

// GetChildGoHomeWithResult represents result wrapper for child go-home-with query.
type GetChildGoHomeWithResult struct {
	Children []ChildGoHomeWithResult `json:"children,omitempty"`
}
