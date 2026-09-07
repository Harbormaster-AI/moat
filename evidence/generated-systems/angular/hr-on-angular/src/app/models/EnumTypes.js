
// enum type JobLevel
export let JobLevel = {
	Entry:"Entry",
	Intermediate:"Intermediate",
	Senior:"Senior",
	Lead:"Lead",
	Manager:"Manager",
	Director:"Director",
	Executive:"Executive",
}

// enum type ExemptStatus
export let ExemptStatus = {
	Exempt:"Exempt",
	NonExempt:"NonExempt",
}

// enum type PositionStatus
export let PositionStatus = {
	Open:"Open",
	Filled:"Filled",
	Frozen:"Frozen",
	Closed:"Closed",
}

// enum type WorkLocationType
export let WorkLocationType = {
	Onsite:"Onsite",
	Hybrid:"Hybrid",
	Remote:"Remote",
}

// enum type EmploymentStatus
export let EmploymentStatus = {
	Active:"Active",
	OnLeave:"OnLeave",
	Suspended:"Suspended",
	Terminated:"Terminated",
}

// enum type AssignmentType
export let AssignmentType = {
	Primary:"Primary",
	Secondary:"Secondary",
	Temporary:"Temporary",
}

// enum type AssignmentStatus
export let AssignmentStatus = {
	Planned:"Planned",
	Active:"Active",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type EmploymentType
export let EmploymentType = {
	FullTime:"FullTime",
	PartTime:"PartTime",
	Temporary:"Temporary",
	Intern:"Intern",
	Contractor:"Contractor",
	Seasonal:"Seasonal",
}

// enum type ContractStatus
export let ContractStatus = {
	Draft:"Draft",
	Active:"Active",
	Suspended:"Suspended",
	Expired:"Expired",
	Terminated:"Terminated",
}

// enum type PayFrequency
export let PayFrequency = {
	Weekly:"Weekly",
	Biweekly:"Biweekly",
	Semimonthly:"Semimonthly",
	Monthly:"Monthly",
	Quarterly:"Quarterly",
}

// enum type ScheduleType
export let ScheduleType = {
	Fixed:"Fixed",
	Flexible:"Flexible",
	Rotating:"Rotating",
}

// enum type DayOfWeek
export let DayOfWeek = {
	Monday:"Monday",
	Tuesday:"Tuesday",
	Wednesday:"Wednesday",
	Thursday:"Thursday",
	Friday:"Friday",
	Saturday:"Saturday",
	Sunday:"Sunday",
}

// enum type SalaryComponentType
export let SalaryComponentType = {
	BaseSalary:"BaseSalary",
	Allowance:"Allowance",
	OvertimeRate:"OvertimeRate",
	Commission:"Commission",
	ShiftDifferential:"ShiftDifferential",
}

// enum type EquityType
export let EquityType = {
	RSU:"RSU",
	StockOption:"StockOption",
	ESPP:"ESPP",
}

// enum type BenefitType
export let BenefitType = {
	Medical:"Medical",
	Dental:"Dental",
	Vision:"Vision",
	LifeInsurance:"LifeInsurance",
	Disability:"Disability",
	Retirement:"Retirement",
	Wellness:"Wellness",
}

// enum type BenefitEnrollmentStatus
export let BenefitEnrollmentStatus = {
	Pending:"Pending",
	Active:"Active",
	Waived:"Waived",
	Cancelled:"Cancelled",
	Terminated:"Terminated",
}

// enum type CoverageLevel
export let CoverageLevel = {
	EmployeeOnly:"EmployeeOnly",
	EmployeeSpouse:"EmployeeSpouse",
	EmployeeChildren:"EmployeeChildren",
	Family:"Family",
}

// enum type DependentRelationship
export let DependentRelationship = {
	Spouse:"Spouse",
	DomesticPartner:"DomesticPartner",
	Child:"Child",
	Other:"Other",
}

// enum type TimeEntryType
export let TimeEntryType = {
	Regular:"Regular",
	Overtime:"Overtime",
	Sick:"Sick",
	Vacation:"Vacation",
	Unpaid:"Unpaid",
}

// enum type TimesheetStatus
export let TimesheetStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	Approved:"Approved",
	Rejected:"Rejected",
	Processed:"Processed",
}

// enum type ApprovalStatus
export let ApprovalStatus = {
	Pending:"Pending",
	Approved:"Approved",
	Rejected:"Rejected",
	Cancelled:"Cancelled",
}

// enum type LeaveCategory
export let LeaveCategory = {
	Vacation:"Vacation",
	Sick:"Sick",
	Parental:"Parental",
	Bereavement:"Bereavement",
	Unpaid:"Unpaid",
	JuryDuty:"JuryDuty",
}

// enum type AccrualUnit
export let AccrualUnit = {
	Hours:"Hours",
	Days:"Days",
}

// enum type LeaveStatus
export let LeaveStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	Approved:"Approved",
	Rejected:"Rejected",
	Cancelled:"Cancelled",
	Taken:"Taken",
}

// enum type PayrollStatus
export let PayrollStatus = {
	Scheduled:"Scheduled",
	InProgress:"InProgress",
	Completed:"Completed",
	Reversed:"Reversed",
}

// enum type PayrollItemType
export let PayrollItemType = {
	Earning:"Earning",
	Deduction:"Deduction",
	Tax:"Tax",
	Benefit:"Benefit",
}

// enum type FilingStatus
export let FilingStatus = {
	Single:"Single",
	MarriedFilingJointly:"MarriedFilingJointly",
	MarriedFilingSeparately:"MarriedFilingSeparately",
	HeadOfHousehold:"HeadOfHousehold",
	QualifyingWidowEr:"QualifyingWidowEr",
}

// enum type CycleStatus
export let CycleStatus = {
	Planned:"Planned",
	Open:"Open",
	Closed:"Closed",
}

// enum type GoalStatus
export let GoalStatus = {
	NotStarted:"NotStarted",
	InProgress:"InProgress",
	Completed:"Completed",
	Deferred:"Deferred",
	Cancelled:"Cancelled",
}

// enum type PerformanceRating
export let PerformanceRating = {
	Unsatisfactory:"Unsatisfactory",
	NeedsImprovement:"NeedsImprovement",
	MeetsExpectations:"MeetsExpectations",
	ExceedsExpectations:"ExceedsExpectations",
	Outstanding:"Outstanding",
}

// enum type ReviewStatus
export let ReviewStatus = {
	NotStarted:"NotStarted",
	InProgress:"InProgress",
	Finalized:"Finalized",
	Acknowledged:"Acknowledged",
}

// enum type DeliveryMethod
export let DeliveryMethod = {
	Classroom:"Classroom",
	Virtual:"Virtual",
	SelfPaced:"SelfPaced",
	Blended:"Blended",
}

// enum type TrainingStatus
export let TrainingStatus = {
	Enrolled:"Enrolled",
	InProgress:"InProgress",
	Completed:"Completed",
	Failed:"Failed",
	Cancelled:"Cancelled",
}

// enum type DisciplinaryActionType
export let DisciplinaryActionType = {
	VerbalWarning:"VerbalWarning",
	WrittenWarning:"WrittenWarning",
	Suspension:"Suspension",
	Termination:"Termination",
}

// enum type RequisitionStatus
export let RequisitionStatus = {
	Draft:"Draft",
	Open:"Open",
	OnHold:"OnHold",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type RequisitionPriority
export let RequisitionPriority = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Critical:"Critical",
}

// enum type CandidateSource
export let CandidateSource = {
	Referral:"Referral",
	Agency:"Agency",
	JobBoard:"JobBoard",
	CareerSite:"CareerSite",
	Campus:"Campus",
	Social:"Social",
	Internal:"Internal",
}

// enum type ApplicationStatus
export let ApplicationStatus = {
	New:"New",
	Screening:"Screening",
	Interview:"Interview",
	Offer:"Offer",
	Hired:"Hired",
	Rejected:"Rejected",
	Withdrawn:"Withdrawn",
}

// enum type InterviewStage
export let InterviewStage = {
	PhoneScreen:"PhoneScreen",
	Technical:"Technical",
	Onsite:"Onsite",
	Panel:"Panel",
	HR:"HR",
	Executive:"Executive",
}

// enum type InterviewResult
export let InterviewResult = {
	Pending:"Pending",
	Proceed:"Proceed",
	Reject:"Reject",
	OfferRecommended:"OfferRecommended",
}

// enum type OfferStatus
export let OfferStatus = {
	Draft:"Draft",
	Sent:"Sent",
	Accepted:"Accepted",
	Declined:"Declined",
	Withdrawn:"Withdrawn",
	Expired:"Expired",
}

// enum type OnboardingTaskStatus
export let OnboardingTaskStatus = {
	NotStarted:"NotStarted",
	InProgress:"InProgress",
	Blocked:"Blocked",
	Completed:"Completed",
}

// enum type BackgroundCheckStatus
export let BackgroundCheckStatus = {
	Ordered:"Ordered",
	InProgress:"InProgress",
	Clear:"Clear",
	Adverse:"Adverse",
	Cancelled:"Cancelled",
}

// enum type DocumentType
export let DocumentType = {
	Resume:"Resume",
	CoverLetter:"CoverLetter",
	ID:"ID",
	Certification:"Certification",
	Contract:"Contract",
	Policy:"Policy",
	Other:"Other",
}

// enum type AcknowledgementStatus
export let AcknowledgementStatus = {
	Pending:"Pending",
	Acknowledged:"Acknowledged",
	Declined:"Declined",
}

// enum type TerminationReason
export let TerminationReason = {
	Voluntary:"Voluntary",
	Involuntary:"Involuntary",
	Retirement:"Retirement",
	Redundancy:"Redundancy",
	EndOfContract:"EndOfContract",
}

// enum type TerminationType
export let TerminationType = {
	Resignation:"Resignation",
	Dismissal:"Dismissal",
	Layoff:"Layoff",
	Retirement:"Retirement",
	EndOfAssignment:"EndOfAssignment",
}

// enum type WorkAuthorizationStatus
export let WorkAuthorizationStatus = {
	NotRequired:"NotRequired",
	Pending:"Pending",
	Authorized:"Authorized",
	Expired:"Expired",
}

// enum type PaymentMethodType
export let PaymentMethodType = {
	DirectDeposit:"DirectDeposit",
	Check:"Check",
	Cash:"Cash",
	InternationalTransfer:"InternationalTransfer",
}
