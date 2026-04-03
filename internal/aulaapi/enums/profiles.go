package enums

// PortalRole is the portal-level role for a user.
type PortalRole string

const (
	PortalRoleOther    PortalRole = "other"
	PortalRoleEmployee PortalRole = "employee"
	PortalRoleChild    PortalRole = "child"
	PortalRoleGuardian PortalRole = "guardian"
	PortalRoleOtp      PortalRole = "otp"
)

// InstitutionRole is the role within an institution.
type InstitutionRole string

const (
	InstitutionRoleUnknown          InstitutionRole = "unknown"
	InstitutionRoleGuardian         InstitutionRole = "guardian"
	InstitutionRoleDaycare          InstitutionRole = "daycare"
	InstitutionRoleLeader           InstitutionRole = "leader"
	InstitutionRolePreschoolTeacher InstitutionRole = "preschoolTeacher"
	InstitutionRoleTeacher          InstitutionRole = "teacher"
	InstitutionRoleEarlyStudent     InstitutionRole = "earlyStudent"
	InstitutionRoleMiddleLateStudent InstitutionRole = "middleLateStudent"
	InstitutionRoleChild            InstitutionRole = "child"
	InstitutionRoleOther            InstitutionRole = "other"
)

// InstitutionTypeEnum is the type of institution.
type InstitutionTypeEnum string

const (
	InstitutionTypeEnumUnknown      InstitutionTypeEnum = "Unknown"
	InstitutionTypeEnumSchool       InstitutionTypeEnum = "School"
	InstitutionTypeEnumDaycare      InstitutionTypeEnum = "Daycare"
	InstitutionTypeEnumMunicipality InstitutionTypeEnum = "Municipality"
	InstitutionTypeEnumCentral      InstitutionTypeEnum = "Central"
)

// GroupRole is the role of a user within a group.
type GroupRole string

const (
	GroupRoleUnknown  GroupRole = "Unknown"
	GroupRoleMember   GroupRole = "Member"
	GroupRoleIndirect GroupRole = "Indirect"
	GroupRoleApplied  GroupRole = "Applied"
	GroupRoleRemoved  GroupRole = "Removed"
	GroupRoleRejected GroupRole = "Rejected"
	GroupRoleInactive GroupRole = "Inactive"
)

// GroupStatus is the status of a group.
type GroupStatus string

const (
	GroupStatusUnidentified GroupStatus = "Unidentified"
	GroupStatusActive       GroupStatus = "Active"
	GroupStatusInactive     GroupStatus = "Inactive"
)

// UserRelationType is the relationship type between users.
type UserRelationType string

const (
	UserRelationTypeOthers   UserRelationType = "Others"
	UserRelationTypeChild    UserRelationType = "Child"
	UserRelationTypeGuardian UserRelationType = "Guardian"
	UserRelationTypeOtp      UserRelationType = "Otp"
	UserRelationTypeTeacher  UserRelationType = "Teacher"
)

// GroupActionType is a group action (join/leave).
type GroupActionType string

const (
	GroupActionTypeLeave GroupActionType = "Leave"
	GroupActionTypeJoin  GroupActionType = "Join"
)

// GroupMembershipRole is the group membership role.
type GroupMembershipRole string

const (
	GroupMembershipRoleOther              GroupMembershipRole = "Other"
	GroupMembershipRoleApplied            GroupMembershipRole = "Applied"
	GroupMembershipRoleMember             GroupMembershipRole = "Member"
	GroupMembershipRoleRemoved            GroupMembershipRole = "Removed"
	GroupMembershipRoleApplicationRemoved GroupMembershipRole = "Application_Removed"
	GroupMembershipRoleIndirect           GroupMembershipRole = "Indirect"
)

// GroupTypeEnum is the type of group.
type GroupTypeEnum string

const (
	GroupTypeEnumUnknown            GroupTypeEnum = "Unknown"
	GroupTypeEnumInstitutional      GroupTypeEnum = "Institutional"
	GroupTypeEnumMunicipal          GroupTypeEnum = "Municipal"
	GroupTypeEnumCrossInstitutional GroupTypeEnum = "Cross_institutional"
	GroupTypeEnumOther              GroupTypeEnum = "Other"
)

// GroupsAccessType is the group access type.
type GroupsAccessType string

const (
	GroupsAccessTypeOther       GroupsAccessType = "Other"
	GroupsAccessTypeClosed      GroupsAccessType = "Closed"
	GroupsAccessTypeOpen        GroupsAccessType = "Open"
	GroupsAccessTypeApplication GroupsAccessType = "Application"
)

// ContactListFilteringProfileType is the contact list filter type.
type ContactListFilteringProfileType string

const (
	ContactListFilteringProfileTypeAllChildren ContactListFilteringProfileType = "AllChildren"
	ContactListFilteringProfileTypeBoy         ContactListFilteringProfileType = "Boy"
	ContactListFilteringProfileTypeGirl        ContactListFilteringProfileType = "Girl"
	ContactListFilteringProfileTypeEmployee    ContactListFilteringProfileType = "Employee"
	ContactListFilteringProfileTypeGuardian    ContactListFilteringProfileType = "Guardian"
)

// ProfileRoleGender is the gender in profile context.
type ProfileRoleGender string

const (
	ProfileRoleGenderBoy     ProfileRoleGender = "Boy"
	ProfileRoleGenderGirl    ProfileRoleGender = "Girl"
	ProfileRoleGenderUnknown ProfileRoleGender = "Unknown"
)

// GetProfileContactSortOrderField is the sort order for contact list.
type GetProfileContactSortOrderField string

const (
	GetProfileContactSortOrderFieldBirthday GetProfileContactSortOrderField = "Birthday"
	GetProfileContactSortOrderFieldName     GetProfileContactSortOrderField = "Name"
)

// GetPersonalReferenceDataOrderField is the sort order for personal reference data.
type GetPersonalReferenceDataOrderField string

const (
	GetPersonalReferenceDataOrderFieldAnswers     GetPersonalReferenceDataOrderField = "Answers"
	GetPersonalReferenceDataOrderFieldDisplayName GetPersonalReferenceDataOrderField = "DisplayName"
)

// LoginAuthenticationMethod is the login authentication method level.
type LoginAuthenticationMethod string

const (
	LoginAuthenticationMethodUnknown        LoginAuthenticationMethod = "Unknown"
	LoginAuthenticationMethodLevel2         LoginAuthenticationMethod = "Level2"
	LoginAuthenticationMethodLevel3NemId    LoginAuthenticationMethod = "Level3NemId"
	LoginAuthenticationMethodLevel3Employees LoginAuthenticationMethod = "Level3Employees"
)

// UpdateProfileInformationReturnCode is the return code from profile update.
type UpdateProfileInformationReturnCode string

const (
	UpdateProfileInfoReturnCodeSuccess                          UpdateProfileInformationReturnCode = "Success"
	UpdateProfileInfoReturnCodeError                            UpdateProfileInformationReturnCode = "Error"
	UpdateProfileInfoReturnCodeErrorUserDeactivated             UpdateProfileInformationReturnCode = "ErrorUserDeactivated"
	UpdateProfileInfoReturnCodeErrorUserAccessDenied            UpdateProfileInformationReturnCode = "ErrorUserAccessDenied"
	UpdateProfileInfoReturnCodeWrongUserTypeLoggedInAsEmployee  UpdateProfileInformationReturnCode = "WrongUserTypeLoggedInAsEmployee"
	UpdateProfileInfoReturnCodeWrongUserTypeLoggedInAsGuardian  UpdateProfileInformationReturnCode = "WrongUserTypeLoggedInAsGuardian"
	UpdateProfileInfoReturnCodeWrongUserTypeLoggedInAsChild     UpdateProfileInformationReturnCode = "WrongUserTypeLoggedInAsChild"
	UpdateProfileInfoReturnCodeWrongUserTypeLoggedInAsOtp       UpdateProfileInformationReturnCode = "WrongUserTypeLoggedInAsOTP"
)

// OnboardingStep is the onboarding step.
type OnboardingStep string

const (
	OnboardingStepAppOnboarding       OnboardingStep = "AppOnboarding"
	OnboardingStepPolicyAcceptance    OnboardingStep = "PolicyAcceptance"
	OnboardingStepMasterData          OnboardingStep = "MasterData"
	OnboardingStepConsents            OnboardingStep = "Consents"
	OnboardingStepAdditionalMasterData OnboardingStep = "AdditionalMasterData"
	OnboardingStepNotificationSettings OnboardingStep = "NotificationSettings"
)
