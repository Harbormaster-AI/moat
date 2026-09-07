package model


//==============================================================
// GovernanceBodyType Declaration
//==============================================================
type GovernanceBodyType int
const (
    GovernanceBodyTypeBoard GovernanceBodyType = iota
	GovernanceBodyTypeCommittee
	GovernanceBodyTypeCouncil
	GovernanceBodyTypeWorkingGroup
)


//==============================================================
// PolicyType Declaration
//==============================================================
type PolicyType int
const (
    PolicyTypeInformationSecurity PolicyType = iota
	PolicyTypeDataProtection
	PolicyTypeEthics
	PolicyTypeRecordsManagement
	PolicyTypeRiskManagement
	PolicyTypeCompliance
	PolicyTypePrivacy
	PolicyTypeAcceptableUse
)


//==============================================================
// DocumentStatus Declaration
//==============================================================
type DocumentStatus int
const (
    DocumentStatusDraft DocumentStatus = iota
	DocumentStatusInReview
	DocumentStatusApproved
	DocumentStatusRetired
)


//==============================================================
// ControlType Declaration
//==============================================================
type ControlType int
const (
    ControlTypePreventive ControlType = iota
	ControlTypeDetective
	ControlTypeCorrective
	ControlTypeDirective
)


//==============================================================
// ControlFrequency Declaration
//==============================================================
type ControlFrequency int
const (
    ControlFrequencyContinuous ControlFrequency = iota
	ControlFrequencyDaily
	ControlFrequencyWeekly
	ControlFrequencyMonthly
	ControlFrequencyQuarterly
	ControlFrequencyAnnually
	ControlFrequencyAdHoc
)


//==============================================================
// ControlStatus Declaration
//==============================================================
type ControlStatus int
const (
    ControlStatusDesigned ControlStatus = iota
	ControlStatusImplemented
	ControlStatusOperating
	ControlStatusRetired
)


//==============================================================
// RiskCategory Declaration
//==============================================================
type RiskCategory int
const (
    RiskCategoryStrategic RiskCategory = iota
	RiskCategoryOperational
	RiskCategoryFinancial
	RiskCategoryCompliance
	RiskCategoryReputational
	RiskCategoryPrivacy
	RiskCategoryCybersecurity
	RiskCategoryThirdParty
)


//==============================================================
// RiskImpact Declaration
//==============================================================
type RiskImpact int
const (
    RiskImpactInsignificant RiskImpact = iota
	RiskImpactMinor
	RiskImpactModerate
	RiskImpactMajor
	RiskImpactSevere
)


//==============================================================
// RiskLikelihood Declaration
//==============================================================
type RiskLikelihood int
const (
    RiskLikelihoodRare RiskLikelihood = iota
	RiskLikelihoodUnlikely
	RiskLikelihoodPossible
	RiskLikelihoodLikely
	RiskLikelihoodAlmostCertain
)


//==============================================================
// RiskStatus Declaration
//==============================================================
type RiskStatus int
const (
    RiskStatusIdentified RiskStatus = iota
	RiskStatusAssessed
	RiskStatusMitigated
	RiskStatusAccepted
	RiskStatusTransferred
	RiskStatusClosed
)


//==============================================================
// AssessmentType Declaration
//==============================================================
type AssessmentType int
const (
    AssessmentTypeSelfAssessment AssessmentType = iota
	AssessmentTypeInternalAssessment
	AssessmentTypeExternalAssessment
	AssessmentTypeReadinessReview
)


//==============================================================
// TestType Declaration
//==============================================================
type TestType int
const (
    TestTypeDesignEffectiveness TestType = iota
	TestTypeOperatingEffectiveness
	TestTypeWalkthrough
	TestTypeReperformance
	TestTypeInquiry
	TestTypeObservation
	TestTypeInspection
	TestTypeDataAnalysis
)


//==============================================================
// ControlEffectiveness Declaration
//==============================================================
type ControlEffectiveness int
const (
    ControlEffectivenessEffective ControlEffectiveness = iota
	ControlEffectivenessPartiallyEffective
	ControlEffectivenessIneffective
	ControlEffectivenessNotTested
)


//==============================================================
// TestStatus Declaration
//==============================================================
type TestStatus int
const (
    TestStatusPlanned TestStatus = iota
	TestStatusInProgress
	TestStatusCompleted
	TestStatusBlocked
	TestStatusCancelled
)


//==============================================================
// EvidenceType Declaration
//==============================================================
type EvidenceType int
const (
    EvidenceTypeDocument EvidenceType = iota
	EvidenceTypeScreenshot
	EvidenceTypeLogExport
	EvidenceTypeSystemReport
	EvidenceTypeTicket
	EvidenceTypeAttestation
	EvidenceTypeConfiguration
	EvidenceTypeDataset
)


//==============================================================
// AuditCycle Declaration
//==============================================================
type AuditCycle int
const (
    AuditCycleAnnual AuditCycle = iota
	AuditCycleSemiAnnual
	AuditCycleQuarterly
	AuditCycleContinuous
	AuditCycleOneTime
)


//==============================================================
// AuditStatus Declaration
//==============================================================
type AuditStatus int
const (
    AuditStatusPlanned AuditStatus = iota
	AuditStatusFieldwork
	AuditStatusReporting
	AuditStatusClosed
	AuditStatusOnHold
)


//==============================================================
// FindingSeverity Declaration
//==============================================================
type FindingSeverity int
const (
    FindingSeverityLow FindingSeverity = iota
	FindingSeverityMedium
	FindingSeverityHigh
	FindingSeverityCritical
)


//==============================================================
// FindingStatus Declaration
//==============================================================
type FindingStatus int
const (
    FindingStatusOpen FindingStatus = iota
	FindingStatusInRemediation
	FindingStatusValidated
	FindingStatusClosed
)


//==============================================================
// ActionStatus Declaration
//==============================================================
type ActionStatus int
const (
    ActionStatusNotStarted ActionStatus = iota
	ActionStatusInProgress
	ActionStatusDeferred
	ActionStatusCompleted
	ActionStatusCancelled
)


//==============================================================
// IssueType Declaration
//==============================================================
type IssueType int
const (
    IssueTypeControlDeficiency IssueType = iota
	IssueTypeProcessGap
	IssueTypeComplianceBreach
	IssueTypeSecurityIncident
	IssueTypeDataQualityIssue
	IssueTypeThirdPartyIssue
)


//==============================================================
// Priority Declaration
//==============================================================
type Priority int
const (
    PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
	PriorityUrgent
)


//==============================================================
// IssueStatus Declaration
//==============================================================
type IssueStatus int
const (
    IssueStatusOpen IssueStatus = iota
	IssueStatusInvestigating
	IssueStatusRemediationPlanned
	IssueStatusRemediationInProgress
	IssueStatusVerified
	IssueStatusClosed
)


//==============================================================
// ComplianceStatus Declaration
//==============================================================
type ComplianceStatus int
const (
    ComplianceStatusNotStarted ComplianceStatus = iota
	ComplianceStatusInProgress
	ComplianceStatusCompliant
	ComplianceStatusNonCompliant
	ComplianceStatusWaived
)


//==============================================================
// Applicability Declaration
//==============================================================
type Applicability int
const (
    ApplicabilityMandatory Applicability = iota
	ApplicabilityRecommended
	ApplicabilityNotApplicable
)


//==============================================================
// AttestationResult Declaration
//==============================================================
type AttestationResult int
const (
    AttestationResultAffirmative AttestationResult = iota
	AttestationResultNegative
	AttestationResultQualified
)


//==============================================================
// LawfulBasis Declaration
//==============================================================
type LawfulBasis int
const (
    LawfulBasisConsent LawfulBasis = iota
	LawfulBasisContract
	LawfulBasisLegalObligation
	LawfulBasisVitalInterests
	LawfulBasisPublicTask
	LawfulBasisLegitimateInterests
)


//==============================================================
// DataClassificationLevel Declaration
//==============================================================
type DataClassificationLevel int
const (
    DataClassificationLevelPublic DataClassificationLevel = iota
	DataClassificationLevelInternal
	DataClassificationLevelConfidential
	DataClassificationLevelRestricted
	DataClassificationLevelHighlyRestricted
)


//==============================================================
// SystemType Declaration
//==============================================================
type SystemType int
const (
    SystemTypeApplication SystemType = iota
	SystemTypeDatabase
	SystemTypeDataWarehouse
	SystemTypeSaaS
	SystemTypeInfrastructure
	SystemTypeEndpoint
)


//==============================================================
// DataSubjectRequestType Declaration
//==============================================================
type DataSubjectRequestType int
const (
    DataSubjectRequestTypeAccess DataSubjectRequestType = iota
	DataSubjectRequestTypeRectification
	DataSubjectRequestTypeErasure
	DataSubjectRequestTypeRestriction
	DataSubjectRequestTypePortability
	DataSubjectRequestTypeObjection
	DataSubjectRequestTypeAutomatedDecisioningReview
)


//==============================================================
// RequestStatus Declaration
//==============================================================
type RequestStatus int
const (
    RequestStatusReceived RequestStatus = iota
	RequestStatusInValidation
	RequestStatusInProgress
	RequestStatusOnHold
	RequestStatusFulfilled
	RequestStatusRejected
)


//==============================================================
// RepositoryType Declaration
//==============================================================
type RepositoryType int
const (
    RepositoryTypeDocumentManagement RepositoryType = iota
	RepositoryTypeRecordsArchive
	RepositoryTypeEmailArchive
	RepositoryTypeFileShare
	RepositoryTypeContentServices
	RepositoryTypeDataLake
)


//==============================================================
// RecordType Declaration
//==============================================================
type RecordType int
const (
    RecordTypePolicyRecord RecordType = iota
	RecordTypeContractRecord
	RecordTypeFinancialRecord
	RecordTypeHRRecord
	RecordTypeCustomerRecord
	RecordTypeTechnicalRecord
	RecordTypeAuditRecord
	RecordTypeLegalRecord
)


//==============================================================
// RecordStatus Declaration
//==============================================================
type RecordStatus int
const (
    RecordStatusActive RecordStatus = iota
	RecordStatusArchived
	RecordStatusPendingDisposition
	RecordStatusDisposed
	RecordStatusOnHold
)


//==============================================================
// RetentionTrigger Declaration
//==============================================================
type RetentionTrigger int
const (
    RetentionTriggerCreationDate RetentionTrigger = iota
	RetentionTriggerLastModified
	RetentionTriggerTermination
	RetentionTriggerContractEnd
	RetentionTriggerEventCompletion
	RetentionTriggerFiscalYearEnd
)


//==============================================================
// DispositionAction Declaration
//==============================================================
type DispositionAction int
const (
    DispositionActionDestroy DispositionAction = iota
	DispositionActionTransferToArchive
	DispositionActionReview
	DispositionActionSecureDelete
	DispositionActionReturnToOwner
)


//==============================================================
// RetentionStatus Declaration
//==============================================================
type RetentionStatus int
const (
    RetentionStatusDraft RetentionStatus = iota
	RetentionStatusApproved
	RetentionStatusInEffect
	RetentionStatusSuspended
	RetentionStatusRetired
)


//==============================================================
// DispositionOutcome Declaration
//==============================================================
type DispositionOutcome int
const (
    DispositionOutcomeApproved DispositionOutcome = iota
	DispositionOutcomeDeferred
	DispositionOutcomeRejected
	DispositionOutcomeExecuted
)


//==============================================================
// LegalHoldStatus Declaration
//==============================================================
type LegalHoldStatus int
const (
    LegalHoldStatusActive LegalHoldStatus = iota
	LegalHoldStatusReleased
	LegalHoldStatusSuperseded
)


//==============================================================
// MatterType Declaration
//==============================================================
type MatterType int
const (
    MatterTypeLitigation MatterType = iota
	MatterTypeInvestigation
	MatterTypeRegulatoryInquiry
	MatterTypeComplaint
	MatterTypeArbitration
)


//==============================================================
// MatterStatus Declaration
//==============================================================
type MatterStatus int
const (
    MatterStatusOpen MatterStatus = iota
	MatterStatusActiveDiscovery
	MatterStatusNegotiation
	MatterStatusSettled
	MatterStatusClosed
)


//==============================================================
// ThirdPartyType Declaration
//==============================================================
type ThirdPartyType int
const (
    ThirdPartyTypeVendor ThirdPartyType = iota
	ThirdPartyTypeProcessor
	ThirdPartyTypeJointController
	ThirdPartyTypeSubprocessor
	ThirdPartyTypePartner
	ThirdPartyTypeConsultant
)


//==============================================================
// VendorCriticality Declaration
//==============================================================
type VendorCriticality int
const (
    VendorCriticalityLow VendorCriticality = iota
	VendorCriticalityMedium
	VendorCriticalityHigh
	VendorCriticalityCritical
)


//==============================================================
// AssessmentResult Declaration
//==============================================================
type AssessmentResult int
const (
    AssessmentResultPass AssessmentResult = iota
	AssessmentResultConditionalPass
	AssessmentResultFail
)


//==============================================================
// ContractStatus Declaration
//==============================================================
type ContractStatus int
const (
    ContractStatusDraft ContractStatus = iota
	ContractStatusActive
	ContractStatusExpiring
	ContractStatusTerminated
	ContractStatusArchived
)


//==============================================================
// ObligationType Declaration
//==============================================================
type ObligationType int
const (
    ObligationTypeRegulatory ObligationType = iota
	ObligationTypeContractual
	ObligationTypePolicyDerived
	ObligationTypeIndustryStandard
)


//==============================================================
// ExceptionType Declaration
//==============================================================
type ExceptionType int
const (
    ExceptionTypePolicyException ExceptionType = iota
	ExceptionTypeControlException
	ExceptionTypeRetentionException
	ExceptionTypeRiskAcceptance
	ExceptionTypeComplianceWaiver
)


//==============================================================
// ExceptionStatus Declaration
//==============================================================
type ExceptionStatus int
const (
    ExceptionStatusDraft ExceptionStatus = iota
	ExceptionStatusSubmitted
	ExceptionStatusApproved
	ExceptionStatusRejected
	ExceptionStatusExpired
)


//==============================================================
// ConsentType Declaration
//==============================================================
type ConsentType int
const (
    ConsentTypeMarketing ConsentType = iota
	ConsentTypeProfiling
	ConsentTypeCookies
	ConsentTypeLocation
	ConsentTypeBiometric
)


//==============================================================
// ConsentStatus Declaration
//==============================================================
type ConsentStatus int
const (
    ConsentStatusGranted ConsentStatus = iota
	ConsentStatusWithdrawn
	ConsentStatusExpired
	ConsentStatusNotRequired
)


//==============================================================
// BreachSeverity Declaration
//==============================================================
type BreachSeverity int
const (
    BreachSeverityLow BreachSeverity = iota
	BreachSeverityMedium
	BreachSeverityHigh
	BreachSeverityCritical
)


//==============================================================
// IncidentStatus Declaration
//==============================================================
type IncidentStatus int
const (
    IncidentStatusIdentified IncidentStatus = iota
	IncidentStatusContained
	IncidentStatusNotified
	IncidentStatusResolved
	IncidentStatusClosed
)

