
// enum type GovernanceBodyType
export let GovernanceBodyType = {
	Board:"Board",
	Committee:"Committee",
	Council:"Council",
	WorkingGroup:"WorkingGroup",
}

// enum type PolicyType
export let PolicyType = {
	InformationSecurity:"InformationSecurity",
	DataProtection:"DataProtection",
	Ethics:"Ethics",
	RecordsManagement:"RecordsManagement",
	RiskManagement:"RiskManagement",
	Compliance:"Compliance",
	Privacy:"Privacy",
	AcceptableUse:"AcceptableUse",
}

// enum type DocumentStatus
export let DocumentStatus = {
	Draft:"Draft",
	InReview:"InReview",
	Approved:"Approved",
	Retired:"Retired",
}

// enum type ControlType
export let ControlType = {
	Preventive:"Preventive",
	Detective:"Detective",
	Corrective:"Corrective",
	Directive:"Directive",
}

// enum type ControlFrequency
export let ControlFrequency = {
	Continuous:"Continuous",
	Daily:"Daily",
	Weekly:"Weekly",
	Monthly:"Monthly",
	Quarterly:"Quarterly",
	Annually:"Annually",
	AdHoc:"AdHoc",
}

// enum type ControlStatus
export let ControlStatus = {
	Designed:"Designed",
	Implemented:"Implemented",
	Operating:"Operating",
	Retired:"Retired",
}

// enum type RiskCategory
export let RiskCategory = {
	Strategic:"Strategic",
	Operational:"Operational",
	Financial:"Financial",
	Compliance:"Compliance",
	Reputational:"Reputational",
	Privacy:"Privacy",
	Cybersecurity:"Cybersecurity",
	ThirdParty:"ThirdParty",
}

// enum type RiskImpact
export let RiskImpact = {
	Insignificant:"Insignificant",
	Minor:"Minor",
	Moderate:"Moderate",
	Major:"Major",
	Severe:"Severe",
}

// enum type RiskLikelihood
export let RiskLikelihood = {
	Rare:"Rare",
	Unlikely:"Unlikely",
	Possible:"Possible",
	Likely:"Likely",
	AlmostCertain:"AlmostCertain",
}

// enum type RiskStatus
export let RiskStatus = {
	Identified:"Identified",
	Assessed:"Assessed",
	Mitigated:"Mitigated",
	Accepted:"Accepted",
	Transferred:"Transferred",
	Closed:"Closed",
}

// enum type AssessmentType
export let AssessmentType = {
	SelfAssessment:"SelfAssessment",
	InternalAssessment:"InternalAssessment",
	ExternalAssessment:"ExternalAssessment",
	ReadinessReview:"ReadinessReview",
}

// enum type TestType
export let TestType = {
	DesignEffectiveness:"DesignEffectiveness",
	OperatingEffectiveness:"OperatingEffectiveness",
	Walkthrough:"Walkthrough",
	Reperformance:"Reperformance",
	Inquiry:"Inquiry",
	Observation:"Observation",
	Inspection:"Inspection",
	DataAnalysis:"DataAnalysis",
}

// enum type ControlEffectiveness
export let ControlEffectiveness = {
	Effective:"Effective",
	PartiallyEffective:"PartiallyEffective",
	Ineffective:"Ineffective",
	NotTested:"NotTested",
}

// enum type TestStatus
export let TestStatus = {
	Planned:"Planned",
	InProgress:"InProgress",
	Completed:"Completed",
	Blocked:"Blocked",
	Cancelled:"Cancelled",
}

// enum type EvidenceType
export let EvidenceType = {
	Document:"Document",
	Screenshot:"Screenshot",
	LogExport:"LogExport",
	SystemReport:"SystemReport",
	Ticket:"Ticket",
	Attestation:"Attestation",
	Configuration:"Configuration",
	Dataset:"Dataset",
}

// enum type AuditCycle
export let AuditCycle = {
	Annual:"Annual",
	SemiAnnual:"SemiAnnual",
	Quarterly:"Quarterly",
	Continuous:"Continuous",
	OneTime:"OneTime",
}

// enum type AuditStatus
export let AuditStatus = {
	Planned:"Planned",
	Fieldwork:"Fieldwork",
	Reporting:"Reporting",
	Closed:"Closed",
	OnHold:"OnHold",
}

// enum type FindingSeverity
export let FindingSeverity = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Critical:"Critical",
}

// enum type FindingStatus
export let FindingStatus = {
	Open:"Open",
	InRemediation:"InRemediation",
	Validated:"Validated",
	Closed:"Closed",
}

// enum type ActionStatus
export let ActionStatus = {
	NotStarted:"NotStarted",
	InProgress:"InProgress",
	Deferred:"Deferred",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type IssueType
export let IssueType = {
	ControlDeficiency:"ControlDeficiency",
	ProcessGap:"ProcessGap",
	ComplianceBreach:"ComplianceBreach",
	SecurityIncident:"SecurityIncident",
	DataQualityIssue:"DataQualityIssue",
	ThirdPartyIssue:"ThirdPartyIssue",
}

// enum type Priority
export let Priority = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Urgent:"Urgent",
}

// enum type IssueStatus
export let IssueStatus = {
	Open:"Open",
	Investigating:"Investigating",
	RemediationPlanned:"RemediationPlanned",
	RemediationInProgress:"RemediationInProgress",
	Verified:"Verified",
	Closed:"Closed",
}

// enum type ComplianceStatus
export let ComplianceStatus = {
	NotStarted:"NotStarted",
	InProgress:"InProgress",
	Compliant:"Compliant",
	NonCompliant:"NonCompliant",
	Waived:"Waived",
}

// enum type Applicability
export let Applicability = {
	Mandatory:"Mandatory",
	Recommended:"Recommended",
	NotApplicable:"NotApplicable",
}

// enum type AttestationResult
export let AttestationResult = {
	Affirmative:"Affirmative",
	Negative:"Negative",
	Qualified:"Qualified",
}

// enum type LawfulBasis
export let LawfulBasis = {
	Consent:"Consent",
	Contract:"Contract",
	LegalObligation:"LegalObligation",
	VitalInterests:"VitalInterests",
	PublicTask:"PublicTask",
	LegitimateInterests:"LegitimateInterests",
}

// enum type DataClassificationLevel
export let DataClassificationLevel = {
	Public:"Public",
	Internal:"Internal",
	Confidential:"Confidential",
	Restricted:"Restricted",
	HighlyRestricted:"HighlyRestricted",
}

// enum type SystemType
export let SystemType = {
	Application:"Application",
	Database:"Database",
	DataWarehouse:"DataWarehouse",
	SaaS:"SaaS",
	Infrastructure:"Infrastructure",
	Endpoint:"Endpoint",
}

// enum type DataSubjectRequestType
export let DataSubjectRequestType = {
	Access:"Access",
	Rectification:"Rectification",
	Erasure:"Erasure",
	Restriction:"Restriction",
	Portability:"Portability",
	Objection:"Objection",
	AutomatedDecisioningReview:"AutomatedDecisioningReview",
}

// enum type RequestStatus
export let RequestStatus = {
	Received:"Received",
	InValidation:"InValidation",
	InProgress:"InProgress",
	OnHold:"OnHold",
	Fulfilled:"Fulfilled",
	Rejected:"Rejected",
}

// enum type RepositoryType
export let RepositoryType = {
	DocumentManagement:"DocumentManagement",
	RecordsArchive:"RecordsArchive",
	EmailArchive:"EmailArchive",
	FileShare:"FileShare",
	ContentServices:"ContentServices",
	DataLake:"DataLake",
}

// enum type RecordType
export let RecordType = {
	PolicyRecord:"PolicyRecord",
	ContractRecord:"ContractRecord",
	FinancialRecord:"FinancialRecord",
	HRRecord:"HRRecord",
	CustomerRecord:"CustomerRecord",
	TechnicalRecord:"TechnicalRecord",
	AuditRecord:"AuditRecord",
	LegalRecord:"LegalRecord",
}

// enum type RecordStatus
export let RecordStatus = {
	Active:"Active",
	Archived:"Archived",
	PendingDisposition:"PendingDisposition",
	Disposed:"Disposed",
	OnHold:"OnHold",
}

// enum type RetentionTrigger
export let RetentionTrigger = {
	CreationDate:"CreationDate",
	LastModified:"LastModified",
	Termination:"Termination",
	ContractEnd:"ContractEnd",
	EventCompletion:"EventCompletion",
	FiscalYearEnd:"FiscalYearEnd",
}

// enum type DispositionAction
export let DispositionAction = {
	Destroy:"Destroy",
	TransferToArchive:"TransferToArchive",
	Review:"Review",
	SecureDelete:"SecureDelete",
	ReturnToOwner:"ReturnToOwner",
}

// enum type RetentionStatus
export let RetentionStatus = {
	Draft:"Draft",
	Approved:"Approved",
	InEffect:"InEffect",
	Suspended:"Suspended",
	Retired:"Retired",
}

// enum type DispositionOutcome
export let DispositionOutcome = {
	Approved:"Approved",
	Deferred:"Deferred",
	Rejected:"Rejected",
	Executed:"Executed",
}

// enum type LegalHoldStatus
export let LegalHoldStatus = {
	Active:"Active",
	Released:"Released",
	Superseded:"Superseded",
}

// enum type MatterType
export let MatterType = {
	Litigation:"Litigation",
	Investigation:"Investigation",
	RegulatoryInquiry:"RegulatoryInquiry",
	Complaint:"Complaint",
	Arbitration:"Arbitration",
}

// enum type MatterStatus
export let MatterStatus = {
	Open:"Open",
	ActiveDiscovery:"ActiveDiscovery",
	Negotiation:"Negotiation",
	Settled:"Settled",
	Closed:"Closed",
}

// enum type ThirdPartyType
export let ThirdPartyType = {
	Vendor:"Vendor",
	Processor:"Processor",
	JointController:"JointController",
	Subprocessor:"Subprocessor",
	Partner:"Partner",
	Consultant:"Consultant",
}

// enum type VendorCriticality
export let VendorCriticality = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Critical:"Critical",
}

// enum type AssessmentResult
export let AssessmentResult = {
	Pass:"Pass",
	ConditionalPass:"ConditionalPass",
	Fail:"Fail",
}

// enum type ContractStatus
export let ContractStatus = {
	Draft:"Draft",
	Active:"Active",
	Expiring:"Expiring",
	Terminated:"Terminated",
	Archived:"Archived",
}

// enum type ObligationType
export let ObligationType = {
	Regulatory:"Regulatory",
	Contractual:"Contractual",
	PolicyDerived:"PolicyDerived",
	IndustryStandard:"IndustryStandard",
}

// enum type ExceptionType
export let ExceptionType = {
	PolicyException:"PolicyException",
	ControlException:"ControlException",
	RetentionException:"RetentionException",
	RiskAcceptance:"RiskAcceptance",
	ComplianceWaiver:"ComplianceWaiver",
}

// enum type ExceptionStatus
export let ExceptionStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	Approved:"Approved",
	Rejected:"Rejected",
	Expired:"Expired",
}

// enum type ConsentType
export let ConsentType = {
	Marketing:"Marketing",
	Profiling:"Profiling",
	Cookies:"Cookies",
	Location:"Location",
	Biometric:"Biometric",
}

// enum type ConsentStatus
export let ConsentStatus = {
	Granted:"Granted",
	Withdrawn:"Withdrawn",
	Expired:"Expired",
	NotRequired:"NotRequired",
}

// enum type BreachSeverity
export let BreachSeverity = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Critical:"Critical",
}

// enum type IncidentStatus
export let IncidentStatus = {
	Identified:"Identified",
	Contained:"Contained",
	Notified:"Notified",
	Resolved:"Resolved",
	Closed:"Closed",
}
