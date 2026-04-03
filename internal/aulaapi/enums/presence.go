package enums

import (
	"encoding/json"
	"fmt"
)

// PresenceTemplateRepeatPattern is the repeat pattern for presence templates.
type PresenceTemplateRepeatPattern string

const (
	PresenceTemplateRepeatNever       PresenceTemplateRepeatPattern = "Never"
	PresenceTemplateRepeatWeekly      PresenceTemplateRepeatPattern = "Weekly"
	PresenceTemplateRepeatEvery2Weeks PresenceTemplateRepeatPattern = "Every2Weeks"
)

// ActivityType is the type of presence activity (SCREAMING_SNAKE_CASE).
type ActivityType string

const (
	ActivityTypePickedUpBy  ActivityType = "PICKED_UP_BY"
	ActivityTypeSelfDecider ActivityType = "SELF_DECIDER"
	ActivityTypeSendHome    ActivityType = "SEND_HOME"
	ActivityTypeGoHomeWith  ActivityType = "GO_HOME_WITH"
	ActivityTypeDropOffTime ActivityType = "DROP_OFF_TIME"
	ActivityTypeSpareTime   ActivityType = "SPARE_TIME"
	ActivityTypeCheckIn     ActivityType = "CHECK_IN"
	ActivityTypeCheckOut    ActivityType = "CHECK_OUT"
	ActivityTypeSleeping    ActivityType = "SLEEPING"
	ActivityTypeAll         ActivityType = "ALL"
)

// ComeGoNotification is the Come and Go notification type (SCREAMING_SNAKE_CASE).
type ComeGoNotification string

const (
	ComeGoNotifAlertResponse    ComeGoNotification = "ALERT_RESPONSE_NOTIFICATION"
	ComeGoNotifAlertInvite      ComeGoNotification = "ALERT_INVITE_NOTIFICATION"
	ComeGoNotifVacationResponse ComeGoNotification = "VACATION_RESPONSE_NOTIFICATION"
)

// ComeGoStaffTab is the staff tab in ComeGo module.
type ComeGoStaffTab string

const (
	ComeGoStaffTabActivityList         ComeGoStaffTab = "ActivityList"
	ComeGoStaffTabWeekOverview         ComeGoStaffTab = "WeekOverview"
	ComeGoStaffTabVacationRegOverview  ComeGoStaffTab = "VacationRegistrationOverview"
	ComeGoStaffTabOpeningHoursAndClosed ComeGoStaffTab = "OpeningHoursAndClosedDays"
)

// ComeGoTab is the user-facing ComeGo tab.
type ComeGoTab string

const (
	ComeGoTabAbsence            ComeGoTab = "AbsenceTab"
	ComeGoTabTime               ComeGoTab = "TimeTab"
	ComeGoTabDailyOverview      ComeGoTab = "DailyOverview"
	ComeGoTabPickupResponsible  ComeGoTab = "PickupResponsible"
	ComeGoTabOpeningHoursInstList ComeGoTab = "OpeningHoursAndClosedDaysInstitutionListPage"
	ComeGoTabOpeningHours       ComeGoTab = "OpeningHoursAndClosedDays"
	ComeGoTabPlanning           ComeGoTab = "PlanningPage"
)

// DepartureType is the departure type for child pickup.
type DepartureType string

const (
	DepartureTypeGoGomeWith          DepartureType = "GoGomeWith"
	DepartureTypeRetrieveResponsible DepartureType = "RetrieveResponsible"
)

// OpeningHoursType is the type of opening hours definition.
type OpeningHoursType string

const (
	OpeningHoursTypeSpecific OpeningHoursType = "SpecificOpeningHours"
	OpeningHoursTypeGeneral  OpeningHoursType = "GeneralOpeningHours"
	OpeningHoursTypeDefault  OpeningHoursType = "DefaultOpeningHours"
	OpeningHoursTypeClosed   OpeningHoursType = "ClosedDay"
)

// PresenceDayOfWeek is the day of week in presence context.
type PresenceDayOfWeek string

const (
	PresenceDayMonday    PresenceDayOfWeek = "Monday"
	PresenceDayTuesday   PresenceDayOfWeek = "Tuesday"
	PresenceDayWednesday PresenceDayOfWeek = "Wednesday"
	PresenceDayThursday  PresenceDayOfWeek = "Thursday"
	PresenceDayFriday    PresenceDayOfWeek = "Friday"
	PresenceDaySaturday  PresenceDayOfWeek = "Saturday"
	PresenceDaySunday    PresenceDayOfWeek = "Sunday"
)

// PresenceModuleSettingsDashboardContext is the dashboard context for presence settings.
type PresenceModuleSettingsDashboardContext string

const (
	PresenceDashboardEmployee PresenceModuleSettingsDashboardContext = "EmployeeDashboardSettings"
	PresenceDashboardCheckin  PresenceModuleSettingsDashboardContext = "CheckinDashboardSettings"
	PresenceDashboardGuardian PresenceModuleSettingsDashboardContext = "GuardianDashboardSettings"
)

// PresenceModuleSettingsModule is the presence module setting type.
type PresenceModuleSettingsModule string

const (
	PresenceModuleDropOffTime       PresenceModuleSettingsModule = "DropOffTime"
	PresenceModuleLocation          PresenceModuleSettingsModule = "Location"
	PresenceModuleSleep             PresenceModuleSettingsModule = "Sleep"
	PresenceModuleFieldTrip         PresenceModuleSettingsModule = "FieldTrip"
	PresenceModulePickupType        PresenceModuleSettingsModule = "PickupType"
	PresenceModulePickupTimes       PresenceModuleSettingsModule = "PickupTimes"
	PresenceModuleDailyMessage      PresenceModuleSettingsModule = "DailyMessage"
	PresenceModuleVacation          PresenceModuleSettingsModule = "Vacation"
	PresenceModuleReportSick        PresenceModuleSettingsModule = "ReportSick"
	PresenceModuleSpareTimeActivity PresenceModuleSettingsModule = "SpareTimeActivity"
)

// PresenceModuleSettingsPermission is the permission level for a presence module setting.
type PresenceModuleSettingsPermission string

const (
	PresencePermissionEditable    PresenceModuleSettingsPermission = "Editable"
	PresencePermissionDeactivated PresenceModuleSettingsPermission = "Deactivated"
	PresencePermissionReadable    PresenceModuleSettingsPermission = "Readable"
)

// PresenceStatus is the presence status of a child.
// The API returns this as an integer (0-10) OR a string name.
type PresenceStatus int

const (
	PresenceStatusNotPresent        PresenceStatus = 0
	PresenceStatusSick              PresenceStatus = 1
	PresenceStatusReportedAbsence   PresenceStatus = 2
	PresenceStatusPresent           PresenceStatus = 3
	PresenceStatusFieldTrip         PresenceStatus = 4
	PresenceStatusSleeping          PresenceStatus = 5
	PresenceStatusSpareTimeActivity PresenceStatus = 6
	PresenceStatusPhysicalPlacement PresenceStatus = 7
	PresenceStatusCheckedOut        PresenceStatus = 8
	PresenceStatusNotArrived        PresenceStatus = 9
	PresenceStatusAll               PresenceStatus = 10
)

var presenceStatusNames = map[string]PresenceStatus{
	"NotPresent":        PresenceStatusNotPresent,
	"Sick":              PresenceStatusSick,
	"ReportedAbsence":   PresenceStatusReportedAbsence,
	"Present":           PresenceStatusPresent,
	"FieldTrip":         PresenceStatusFieldTrip,
	"Sleeping":          PresenceStatusSleeping,
	"SpareTimeActivity": PresenceStatusSpareTimeActivity,
	"PhysicalPlacement": PresenceStatusPhysicalPlacement,
	"CheckedOut":        PresenceStatusCheckedOut,
	"NotArrived":        PresenceStatusNotArrived,
	"All":               PresenceStatusAll,
}

// UnmarshalJSON accepts both integer (0-10) and string representations.
func (p *PresenceStatus) UnmarshalJSON(data []byte) error {
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		if num < 0 || num > 10 {
			return fmt.Errorf("unknown presence status: %d", num)
		}
		*p = PresenceStatus(num)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if v, ok := presenceStatusNames[s]; ok {
			*p = v
			return nil
		}
		return fmt.Errorf("unknown presence status: %s", s)
	}
	return fmt.Errorf("PresenceStatus: cannot unmarshal %s", string(data))
}

// MarshalJSON serializes as integer.
func (p PresenceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(p))
}

// String returns the name of the presence status.
func (p PresenceStatus) String() string {
	for name, val := range presenceStatusNames {
		if val == p {
			return name
		}
	}
	return fmt.Sprintf("PresenceStatus(%d)", int(p))
}

// PresenceTemplateEditingOption is the editing option for presence templates.
type PresenceTemplateEditingOption string

const (
	PresenceTemplateEditSingleDay     PresenceTemplateEditingOption = "EditSingleDay"
	PresenceTemplateEditWholeTemplate PresenceTemplateEditingOption = "EditWholeTemplate"
	PresenceTemplateEditDelete        PresenceTemplateEditingOption = "Delete"
)

// SpareTimeActivityAction is the action on a spare time activity.
type SpareTimeActivityAction string

const (
	SpareTimeActivityActionEdit   SpareTimeActivityAction = "Edit"
	SpareTimeActivityActionDelete SpareTimeActivityAction = "Delete"
)

// TokenStatus is the token usage status.
type TokenStatus string

const (
	TokenStatusUsed    TokenStatus = "Used"
	TokenStatusNotUsed TokenStatus = "NotUsed"
	TokenStatusExpired TokenStatus = "Expired"
)

// ComeGoEmployeeWeekOverviewTense is the tense of employee week overview (camelCase).
type ComeGoEmployeeWeekOverviewTense string

const (
	ComeGoWeekTensePast            ComeGoEmployeeWeekOverviewTense = "past"
	ComeGoWeekTensePresent         ComeGoEmployeeWeekOverviewTense = "present"
	ComeGoWeekTenseNotSpecified    ComeGoEmployeeWeekOverviewTense = "notSpecified"
	ComeGoWeekTenseMissingCheckout ComeGoEmployeeWeekOverviewTense = "missingCheckout"
)

// ComeGoEmployeeWeekOverviewFilterOption is the filter option for employee week overview.
type ComeGoEmployeeWeekOverviewFilterOption string

const (
	ComeGoWeekFilterPresent    ComeGoEmployeeWeekOverviewFilterOption = "Present"
	ComeGoWeekFilterVacation   ComeGoEmployeeWeekOverviewFilterOption = "Vacation"
	ComeGoWeekFilterSick       ComeGoEmployeeWeekOverviewFilterOption = "Sick"
	ComeGoWeekFilterNotArrived ComeGoEmployeeWeekOverviewFilterOption = "NotArrived"
)

// ComeGoEmployeeWeekOverviewPresenceType is the presence type in employee week overview.
type ComeGoEmployeeWeekOverviewPresenceType string

const (
	ComeGoWeekPresencePresent    ComeGoEmployeeWeekOverviewPresenceType = "Present"
	ComeGoWeekPresenceVacation   ComeGoEmployeeWeekOverviewPresenceType = "Vacation"
	ComeGoWeekPresenceSick       ComeGoEmployeeWeekOverviewPresenceType = "Sick"
	ComeGoWeekPresenceNotArrived ComeGoEmployeeWeekOverviewPresenceType = "NotArrived"
	ComeGoWeekPresenceNone       ComeGoEmployeeWeekOverviewPresenceType = "None"
)

// ComeGoType is the ComeGo type for remote notifications.
type ComeGoType string

const (
	ComeGoTypePickupActivity  ComeGoType = "PickupActivity"
	ComeGoTypeVacationRequest ComeGoType = "VacationRegistrationRequest"
)
