package model


//==============================================================
// JobLevel Declaration
//==============================================================
type JobLevel int
const (
    JobLevelEntry JobLevel = iota
	JobLevelIntermediate
	JobLevelSenior
	JobLevelLead
	JobLevelManager
	JobLevelDirector
	JobLevelExecutive
)


//==============================================================
// ExemptStatus Declaration
//==============================================================
type ExemptStatus int
const (
    ExemptStatusExempt ExemptStatus = iota
	ExemptStatusNonExempt
)


//==============================================================
// PositionStatus Declaration
//==============================================================
type PositionStatus int
const (
    PositionStatusOpen PositionStatus = iota
	PositionStatusFilled
	PositionStatusFrozen
	PositionStatusClosed
)


//==============================================================
// WorkLocationType Declaration
//==============================================================
type WorkLocationType int
const (
    WorkLocationTypeOnsite WorkLocationType = iota
	WorkLocationTypeHybrid
	WorkLocationTypeRemote
)


//==============================================================
// EmploymentStatus Declaration
//==============================================================
type EmploymentStatus int
const (
    EmploymentStatusActive EmploymentStatus = iota
	EmploymentStatusOnLeave
	EmploymentStatusSuspended
	EmploymentStatusTerminated
)


//==============================================================
// AssignmentType Declaration
//==============================================================
type AssignmentType int
const (
    AssignmentTypePrimary AssignmentType = iota
	AssignmentTypeSecondary
	AssignmentTypeTemporary
)


//==============================================================
// AssignmentStatus Declaration
//==============================================================
type AssignmentStatus int
const (
    AssignmentStatusPlanned AssignmentStatus = iota
	AssignmentStatusActive
	AssignmentStatusCompleted
	AssignmentStatusCancelled
)


//==============================================================
// EmploymentType Declaration
//==============================================================
type EmploymentType int
const (
    EmploymentTypeFullTime EmploymentType = iota
	EmploymentTypePartTime
	EmploymentTypeTemporary
	EmploymentTypeIntern
	EmploymentTypeContractor
	EmploymentTypeSeasonal
)


//==============================================================
// ContractStatus Declaration
//==============================================================
type ContractStatus int
const (
    ContractStatusDraft ContractStatus = iota
	ContractStatusActive
	ContractStatusSuspended
	ContractStatusExpired
	ContractStatusTerminated
)


//==============================================================
// PayFrequency Declaration
//==============================================================
type PayFrequency int
const (
    PayFrequencyWeekly PayFrequency = iota
	PayFrequencyBiweekly
	PayFrequencySemimonthly
	PayFrequencyMonthly
	PayFrequencyQuarterly
)


//==============================================================
// ScheduleType Declaration
//==============================================================
type ScheduleType int
const (
    ScheduleTypeFixed ScheduleType = iota
	ScheduleTypeFlexible
	ScheduleTypeRotating
)


//==============================================================
// DayOfWeek Declaration
//==============================================================
type DayOfWeek int
const (
    DayOfWeekMonday DayOfWeek = iota
	DayOfWeekTuesday
	DayOfWeekWednesday
	DayOfWeekThursday
	DayOfWeekFriday
	DayOfWeekSaturday
	DayOfWeekSunday
)


//==============================================================
// SalaryComponentType Declaration
//==============================================================
type SalaryComponentType int
const (
    SalaryComponentTypeBaseSalary SalaryComponentType = iota
	SalaryComponentTypeAllowance
	SalaryComponentTypeOvertimeRate
	SalaryComponentTypeCommission
	SalaryComponentTypeShiftDifferential
)


//==============================================================
// EquityType Declaration
//==============================================================
type EquityType int
const (
    EquityTypeRSU EquityType = iota
	EquityTypeStockOption
	EquityTypeESPP
)


//==============================================================
// BenefitType Declaration
//==============================================================
type BenefitType int
const (
    BenefitTypeMedical BenefitType = iota
	BenefitTypeDental
	BenefitTypeVision
	BenefitTypeLifeInsurance
	BenefitTypeDisability
	BenefitTypeRetirement
	BenefitTypeWellness
)


//==============================================================
// BenefitEnrollmentStatus Declaration
//==============================================================
type BenefitEnrollmentStatus int
const (
    BenefitEnrollmentStatusPending BenefitEnrollmentStatus = iota
	BenefitEnrollmentStatusActive
	BenefitEnrollmentStatusWaived
	BenefitEnrollmentStatusCancelled
	BenefitEnrollmentStatusTerminated
)


//==============================================================
// CoverageLevel Declaration
//==============================================================
type CoverageLevel int
const (
    CoverageLevelEmployeeOnly CoverageLevel = iota
	CoverageLevelEmployeeSpouse
	CoverageLevelEmployeeChildren
	CoverageLevelFamily
)


//==============================================================
// DependentRelationship Declaration
//==============================================================
type DependentRelationship int
const (
    DependentRelationshipSpouse DependentRelationship = iota
	DependentRelationshipDomesticPartner
	DependentRelationshipChild
	DependentRelationshipOther
)


//==============================================================
// TimeEntryType Declaration
//==============================================================
type TimeEntryType int
const (
    TimeEntryTypeRegular TimeEntryType = iota
	TimeEntryTypeOvertime
	TimeEntryTypeSick
	TimeEntryTypeVacation
	TimeEntryTypeUnpaid
)


//==============================================================
// TimesheetStatus Declaration
//==============================================================
type TimesheetStatus int
const (
    TimesheetStatusDraft TimesheetStatus = iota
	TimesheetStatusSubmitted
	TimesheetStatusApproved
	TimesheetStatusRejected
	TimesheetStatusProcessed
)


//==============================================================
// ApprovalStatus Declaration
//==============================================================
type ApprovalStatus int
const (
    ApprovalStatusPending ApprovalStatus = iota
	ApprovalStatusApproved
	ApprovalStatusRejected
	ApprovalStatusCancelled
)


//==============================================================
// LeaveCategory Declaration
//==============================================================
type LeaveCategory int
const (
    LeaveCategoryVacation LeaveCategory = iota
	LeaveCategorySick
	LeaveCategoryParental
	LeaveCategoryBereavement
	LeaveCategoryUnpaid
	LeaveCategoryJuryDuty
)


//==============================================================
// AccrualUnit Declaration
//==============================================================
type AccrualUnit int
const (
    AccrualUnitHours AccrualUnit = iota
	AccrualUnitDays
)


//==============================================================
// LeaveStatus Declaration
//==============================================================
type LeaveStatus int
const (
    LeaveStatusDraft LeaveStatus = iota
	LeaveStatusSubmitted
	LeaveStatusApproved
	LeaveStatusRejected
	LeaveStatusCancelled
	LeaveStatusTaken
)


//==============================================================
// PayrollStatus Declaration
//==============================================================
type PayrollStatus int
const (
    PayrollStatusScheduled PayrollStatus = iota
	PayrollStatusInProgress
	PayrollStatusCompleted
	PayrollStatusReversed
)


//==============================================================
// PayrollItemType Declaration
//==============================================================
type PayrollItemType int
const (
    PayrollItemTypeEarning PayrollItemType = iota
	PayrollItemTypeDeduction
	PayrollItemTypeTax
	PayrollItemTypeBenefit
)


//==============================================================
// FilingStatus Declaration
//==============================================================
type FilingStatus int
const (
    FilingStatusSingle FilingStatus = iota
	FilingStatusMarriedFilingJointly
	FilingStatusMarriedFilingSeparately
	FilingStatusHeadOfHousehold
	FilingStatusQualifyingWidowEr
)


//==============================================================
// CycleStatus Declaration
//==============================================================
type CycleStatus int
const (
    CycleStatusPlanned CycleStatus = iota
	CycleStatusOpen
	CycleStatusClosed
)


//==============================================================
// GoalStatus Declaration
//==============================================================
type GoalStatus int
const (
    GoalStatusNotStarted GoalStatus = iota
	GoalStatusInProgress
	GoalStatusCompleted
	GoalStatusDeferred
	GoalStatusCancelled
)


//==============================================================
// PerformanceRating Declaration
//==============================================================
type PerformanceRating int
const (
    PerformanceRatingUnsatisfactory PerformanceRating = iota
	PerformanceRatingNeedsImprovement
	PerformanceRatingMeetsExpectations
	PerformanceRatingExceedsExpectations
	PerformanceRatingOutstanding
)


//==============================================================
// ReviewStatus Declaration
//==============================================================
type ReviewStatus int
const (
    ReviewStatusNotStarted ReviewStatus = iota
	ReviewStatusInProgress
	ReviewStatusFinalized
	ReviewStatusAcknowledged
)


//==============================================================
// DeliveryMethod Declaration
//==============================================================
type DeliveryMethod int
const (
    DeliveryMethodClassroom DeliveryMethod = iota
	DeliveryMethodVirtual
	DeliveryMethodSelfPaced
	DeliveryMethodBlended
)


//==============================================================
// TrainingStatus Declaration
//==============================================================
type TrainingStatus int
const (
    TrainingStatusEnrolled TrainingStatus = iota
	TrainingStatusInProgress
	TrainingStatusCompleted
	TrainingStatusFailed
	TrainingStatusCancelled
)


//==============================================================
// DisciplinaryActionType Declaration
//==============================================================
type DisciplinaryActionType int
const (
    DisciplinaryActionTypeVerbalWarning DisciplinaryActionType = iota
	DisciplinaryActionTypeWrittenWarning
	DisciplinaryActionTypeSuspension
	DisciplinaryActionTypeTermination
)


//==============================================================
// RequisitionStatus Declaration
//==============================================================
type RequisitionStatus int
const (
    RequisitionStatusDraft RequisitionStatus = iota
	RequisitionStatusOpen
	RequisitionStatusOnHold
	RequisitionStatusClosed
	RequisitionStatusCancelled
)


//==============================================================
// RequisitionPriority Declaration
//==============================================================
type RequisitionPriority int
const (
    RequisitionPriorityLow RequisitionPriority = iota
	RequisitionPriorityMedium
	RequisitionPriorityHigh
	RequisitionPriorityCritical
)


//==============================================================
// CandidateSource Declaration
//==============================================================
type CandidateSource int
const (
    CandidateSourceReferral CandidateSource = iota
	CandidateSourceAgency
	CandidateSourceJobBoard
	CandidateSourceCareerSite
	CandidateSourceCampus
	CandidateSourceSocial
	CandidateSourceInternal
)


//==============================================================
// ApplicationStatus Declaration
//==============================================================
type ApplicationStatus int
const (
    ApplicationStatusNew ApplicationStatus = iota
	ApplicationStatusScreening
	ApplicationStatusInterview
	ApplicationStatusOffer
	ApplicationStatusHired
	ApplicationStatusRejected
	ApplicationStatusWithdrawn
)


//==============================================================
// InterviewStage Declaration
//==============================================================
type InterviewStage int
const (
    InterviewStagePhoneScreen InterviewStage = iota
	InterviewStageTechnical
	InterviewStageOnsite
	InterviewStagePanel
	InterviewStageHR
	InterviewStageExecutive
)


//==============================================================
// InterviewResult Declaration
//==============================================================
type InterviewResult int
const (
    InterviewResultPending InterviewResult = iota
	InterviewResultProceed
	InterviewResultReject
	InterviewResultOfferRecommended
)


//==============================================================
// OfferStatus Declaration
//==============================================================
type OfferStatus int
const (
    OfferStatusDraft OfferStatus = iota
	OfferStatusSent
	OfferStatusAccepted
	OfferStatusDeclined
	OfferStatusWithdrawn
	OfferStatusExpired
)


//==============================================================
// OnboardingTaskStatus Declaration
//==============================================================
type OnboardingTaskStatus int
const (
    OnboardingTaskStatusNotStarted OnboardingTaskStatus = iota
	OnboardingTaskStatusInProgress
	OnboardingTaskStatusBlocked
	OnboardingTaskStatusCompleted
)


//==============================================================
// BackgroundCheckStatus Declaration
//==============================================================
type BackgroundCheckStatus int
const (
    BackgroundCheckStatusOrdered BackgroundCheckStatus = iota
	BackgroundCheckStatusInProgress
	BackgroundCheckStatusClear
	BackgroundCheckStatusAdverse
	BackgroundCheckStatusCancelled
)


//==============================================================
// DocumentType Declaration
//==============================================================
type DocumentType int
const (
    DocumentTypeResume DocumentType = iota
	DocumentTypeCoverLetter
	DocumentTypeID
	DocumentTypeCertification
	DocumentTypeContract
	DocumentTypePolicy
	DocumentTypeOther
)


//==============================================================
// AcknowledgementStatus Declaration
//==============================================================
type AcknowledgementStatus int
const (
    AcknowledgementStatusPending AcknowledgementStatus = iota
	AcknowledgementStatusAcknowledged
	AcknowledgementStatusDeclined
)


//==============================================================
// TerminationReason Declaration
//==============================================================
type TerminationReason int
const (
    TerminationReasonVoluntary TerminationReason = iota
	TerminationReasonInvoluntary
	TerminationReasonRetirement
	TerminationReasonRedundancy
	TerminationReasonEndOfContract
)


//==============================================================
// TerminationType Declaration
//==============================================================
type TerminationType int
const (
    TerminationTypeResignation TerminationType = iota
	TerminationTypeDismissal
	TerminationTypeLayoff
	TerminationTypeRetirement
	TerminationTypeEndOfAssignment
)


//==============================================================
// WorkAuthorizationStatus Declaration
//==============================================================
type WorkAuthorizationStatus int
const (
    WorkAuthorizationStatusNotRequired WorkAuthorizationStatus = iota
	WorkAuthorizationStatusPending
	WorkAuthorizationStatusAuthorized
	WorkAuthorizationStatusExpired
)


//==============================================================
// PaymentMethodType Declaration
//==============================================================
type PaymentMethodType int
const (
    PaymentMethodTypeDirectDeposit PaymentMethodType = iota
	PaymentMethodTypeCheck
	PaymentMethodTypeCash
	PaymentMethodTypeInternationalTransfer
)

