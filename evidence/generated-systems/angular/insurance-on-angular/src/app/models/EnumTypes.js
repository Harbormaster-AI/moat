
// enum type LineOfBusiness
export let LineOfBusiness = {
	PersonalAuto:"PersonalAuto",
	Homeowners:"Homeowners",
	Renters:"Renters",
	TermLife:"TermLife",
	WholeLife:"WholeLife",
	CommercialProperty:"CommercialProperty",
	GeneralLiability:"GeneralLiability",
	WorkersCompensation:"WorkersCompensation",
}

// enum type CoverageType
export let CoverageType = {
	Liability:"Liability",
	Collision:"Collision",
	Comprehensive:"Comprehensive",
	PropertyDamage:"PropertyDamage",
	BodilyInjury:"BodilyInjury",
	UninsuredMotorist:"UninsuredMotorist",
	MedicalPayments:"MedicalPayments",
	Dwelling:"Dwelling",
	Contents:"Contents",
	PersonalLiability:"PersonalLiability",
	BusinessInterruption:"BusinessInterruption",
	ProfessionalLiability:"ProfessionalLiability",
}

// enum type DistributionChannelType
export let DistributionChannelType = {
	Agency:"Agency",
	Broker:"Broker",
	Direct:"Direct",
	Bancassurance:"Bancassurance",
	AffinityPartner:"AffinityPartner",
	OnlineAggregator:"OnlineAggregator",
}

// enum type ProducerStatus
export let ProducerStatus = {
	Active:"Active",
	Suspended:"Suspended",
	Terminated:"Terminated",
}

// enum type CustomerType
export let CustomerType = {
	Individual:"Individual",
	Organization:"Organization",
}

// enum type ApplicationStatus
export let ApplicationStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	UnderReview:"UnderReview",
	Quoted:"Quoted",
	Declined:"Declined",
	Withdrawn:"Withdrawn",
	Bound:"Bound",
}

// enum type UnderwritingDecisionType
export let UnderwritingDecisionType = {
	Approve:"Approve",
	ConditionalApprove:"ConditionalApprove",
	Refer:"Refer",
	Decline:"Decline",
}

// enum type PolicyStatus
export let PolicyStatus = {
	Quoted:"Quoted",
	Active:"Active",
	Lapsed:"Lapsed",
	Cancelled:"Cancelled",
	Expired:"Expired",
	PendingCancel:"PendingCancel",
	PendingReinstatement:"PendingReinstatement",
}

// enum type PaymentPlanType
export let PaymentPlanType = {
	Annual:"Annual",
	SemiAnnual:"SemiAnnual",
	Quarterly:"Quarterly",
	Monthly:"Monthly",
	PayInFull:"PayInFull",
}

// enum type BillingStatus
export let BillingStatus = {
	Current:"Current",
	Delinquent:"Delinquent",
	Collections:"Collections",
	Closed:"Closed",
}

// enum type InvoiceStatus
export let InvoiceStatus = {
	Open:"Open",
	Paid:"Paid",
	PartiallyPaid:"PartiallyPaid",
	Void:"Void",
}

// enum type PaymentMethod
export let PaymentMethod = {
	ACH:"ACH",
	CreditCard:"CreditCard",
	DebitCard:"DebitCard",
	Check:"Check",
	Cash:"Cash",
	Wire:"Wire",
}

// enum type PaymentStatus
export let PaymentStatus = {
	Pending:"Pending",
	Settled:"Settled",
	Failed:"Failed",
	Refunded:"Refunded",
	Reversed:"Reversed",
}

// enum type InsuredObjectType
export let InsuredObjectType = {
	Vehicle:"Vehicle",
	Property:"Property",
	Person:"Person",
	Equipment:"Equipment",
	LiabilityExposure:"LiabilityExposure",
}

// enum type RelationshipType
export let RelationshipType = {
	Spouse:"Spouse",
	Child:"Child",
	Parent:"Parent",
	Sibling:"Sibling",
	BusinessPartner:"BusinessPartner",
	Estate:"Estate",
	Trust:"Trust",
	Other:"Other",
}

// enum type ClaimStatus
export let ClaimStatus = {
	Open:"Open",
	Closed:"Closed",
	Reopened:"Reopened",
	Denied:"Denied",
	PendingInvestigation:"PendingInvestigation",
	Litigation:"Litigation",
}

// enum type PerilType
export let PerilType = {
	AutoAccident:"AutoAccident",
	Fire:"Fire",
	Theft:"Theft",
	Windstorm:"Windstorm",
	Flood:"Flood",
	Hail:"Hail",
	Earthquake:"Earthquake",
	Vandalism:"Vandalism",
	Injury:"Injury",
	Death:"Death",
}

// enum type CauseOfLoss
export let CauseOfLoss = {
	Collision:"Collision",
	Weather:"Weather",
	MechanicalFailure:"MechanicalFailure",
	HumanError:"HumanError",
	NaturalDisaster:"NaturalDisaster",
	Theft:"Theft",
	Vandalism:"Vandalism",
	LiabilityClaim:"LiabilityClaim",
	Illness:"Illness",
}

// enum type ExposureType
export let ExposureType = {
	BodilyInjury:"BodilyInjury",
	PropertyDamage:"PropertyDamage",
	Medical:"Medical",
	UninsuredMotorist:"UninsuredMotorist",
	PersonalInjuryProtection:"PersonalInjuryProtection",
	DwellingDamage:"DwellingDamage",
	ContentsDamage:"ContentsDamage",
	BusinessIncome:"BusinessIncome",
}

// enum type ExposureStatus
export let ExposureStatus = {
	Open:"Open",
	Closed:"Closed",
	Pending:"Pending",
	Reserved:"Reserved",
}

// enum type AdjusterType
export let AdjusterType = {
	Staff:"Staff",
	Independent:"Independent",
	Public:"Public",
}

// enum type ReserveType
export let ReserveType = {
	Indemnity:"Indemnity",
	Expense:"Expense",
	Legal:"Legal",
	Medical:"Medical",
}

// enum type ReserveStatus
export let ReserveStatus = {
	Open:"Open",
	Released:"Released",
	Increased:"Increased",
	Decreased:"Decreased",
	Closed:"Closed",
}

// enum type PayeeType
export let PayeeType = {
	Claimant:"Claimant",
	Beneficiary:"Beneficiary",
	ServiceProvider:"ServiceProvider",
	Lienholder:"Lienholder",
	Attorney:"Attorney",
}

// enum type ServiceProviderType
export let ServiceProviderType = {
	RepairShop:"RepairShop",
	Towing:"Towing",
	MedicalProvider:"MedicalProvider",
	Attorney:"Attorney",
	ForensicEngineer:"ForensicEngineer",
	RentalCar:"RentalCar",
}

// enum type NetworkStatus
export let NetworkStatus = {
	InNetwork:"InNetwork",
	OutOfNetwork:"OutOfNetwork",
}

// enum type ReinsuranceType
export let ReinsuranceType = {
	Treaty:"Treaty",
	Facultative:"Facultative",
}

// enum type TreatyType
export let TreatyType = {
	QuotaShare:"QuotaShare",
	Surplus:"Surplus",
	ExcessOfLoss:"ExcessOfLoss",
	StopLoss:"StopLoss",
}

// enum type ThirdPartyType
export let ThirdPartyType = {
	Individual:"Individual",
	Company:"Company",
	GovernmentAgency:"GovernmentAgency",
}

// enum type DocumentType
export let DocumentType = {
	ApplicationForm:"ApplicationForm",
	PolicyDocument:"PolicyDocument",
	Endorsement:"Endorsement",
	Invoice:"Invoice",
	ClaimForm:"ClaimForm",
	PoliceReport:"PoliceReport",
	Estimate:"Estimate",
	Photo:"Photo",
	MedicalRecord:"MedicalRecord",
	Correspondence:"Correspondence",
}

// enum type SubrogationStatus
export let SubrogationStatus = {
	Open:"Open",
	Negotiating:"Negotiating",
	Settled:"Settled",
	Uncollectible:"Uncollectible",
	Closed:"Closed",
}
