package model


//==============================================================
// UserRole Declaration
//==============================================================
type UserRole int
const (
    UserRoleSalesRep UserRole = iota
	UserRoleSalesManager
	UserRoleServiceAgent
	UserRoleMarketingSpecialist
	UserRoleAdministrator
	UserRoleExecutive
)


//==============================================================
// UserStatus Declaration
//==============================================================
type UserStatus int
const (
    UserStatusActive UserStatus = iota
	UserStatusInactive
	UserStatusLocked
	UserStatusPendingInvite
)


//==============================================================
// TeamType Declaration
//==============================================================
type TeamType int
const (
    TeamTypeSales TeamType = iota
	TeamTypeService
	TeamTypeMarketing
	TeamTypeAccountTeam
	TeamTypeDealDesk
	TeamTypeCrossFunctional
)


//==============================================================
// TerritoryType Declaration
//==============================================================
type TerritoryType int
const (
    TerritoryTypeGeographic TerritoryType = iota
	TerritoryTypeIndustry
	TerritoryTypeNamedAccount
	TerritoryTypeSegment
	TerritoryTypeHybrid
)


//==============================================================
// AccountType Declaration
//==============================================================
type AccountType int
const (
    AccountTypeProspect AccountType = iota
	AccountTypeCustomer
	AccountTypePartner
	AccountTypeVendor
	AccountTypeCompetitor
)


//==============================================================
// AccountLifecycleStage Declaration
//==============================================================
type AccountLifecycleStage int
const (
    AccountLifecycleStageSubscriber AccountLifecycleStage = iota
	AccountLifecycleStageLead
	AccountLifecycleStageMarketingQualified
	AccountLifecycleStageSalesQualified
	AccountLifecycleStageCustomer
	AccountLifecycleStageEvangelist
	AccountLifecycleStageChurned
)


//==============================================================
// ContactMethod Declaration
//==============================================================
type ContactMethod int
const (
    ContactMethodEmail ContactMethod = iota
	ContactMethodPhone
	ContactMethodMobile
	ContactMethodSMS
	ContactMethodInPerson
	ContactMethodWeb
)


//==============================================================
// LeadStatus Declaration
//==============================================================
type LeadStatus int
const (
    LeadStatusNew LeadStatus = iota
	LeadStatusWorking
	LeadStatusNurturing
	LeadStatusQualified
	LeadStatusUnqualified
	LeadStatusConverted
)


//==============================================================
// LeadSource Declaration
//==============================================================
type LeadSource int
const (
    LeadSourceWeb LeadSource = iota
	LeadSourceReferral
	LeadSourceEvent
	LeadSourcePartner
	LeadSourceAdvertisement
	LeadSourceOutbound
	LeadSourceInbound
	LeadSourceSocial
	LeadSourceOther
)


//==============================================================
// LeadRating Declaration
//==============================================================
type LeadRating int
const (
    LeadRatingHot LeadRating = iota
	LeadRatingWarm
	LeadRatingCold
)


//==============================================================
// OpportunityStage Declaration
//==============================================================
type OpportunityStage int
const (
    OpportunityStageQualification OpportunityStage = iota
	OpportunityStageDiscovery
	OpportunityStageProposal
	OpportunityStageNegotiation
	OpportunityStageClosedWon
	OpportunityStageClosedLost
)


//==============================================================
// OpportunityType Declaration
//==============================================================
type OpportunityType int
const (
    OpportunityTypeNewBusiness OpportunityType = iota
	OpportunityTypeExistingBusiness
	OpportunityTypeRenewal
	OpportunityTypeUpsell
	OpportunityTypeCrossSell
)


//==============================================================
// ForecastCategory Declaration
//==============================================================
type ForecastCategory int
const (
    ForecastCategoryPipeline ForecastCategory = iota
	ForecastCategoryBestCase
	ForecastCategoryCommit
	ForecastCategoryOmitted
	ForecastCategoryClosed
)


//==============================================================
// ProductType Declaration
//==============================================================
type ProductType int
const (
    ProductTypeGood ProductType = iota
	ProductTypeService
	ProductTypeSubscription
	ProductTypeBundle
)


//==============================================================
// UnitOfMeasure Declaration
//==============================================================
type UnitOfMeasure int
const (
    UnitOfMeasureEach UnitOfMeasure = iota
	UnitOfMeasureHour
	UnitOfMeasureDay
	UnitOfMeasureMonth
	UnitOfMeasureUser
	UnitOfMeasurePackage
)


//==============================================================
// QuoteStatus Declaration
//==============================================================
type QuoteStatus int
const (
    QuoteStatusDraft QuoteStatus = iota
	QuoteStatusPresented
	QuoteStatusApproved
	QuoteStatusRejected
	QuoteStatusAccepted
	QuoteStatusExpired
	QuoteStatusWithdrawn
)


//==============================================================
// OrderStatus Declaration
//==============================================================
type OrderStatus int
const (
    OrderStatusDraft OrderStatus = iota
	OrderStatusSubmitted
	OrderStatusPartiallyFulfilled
	OrderStatusFulfilled
	OrderStatusInvoiced
	OrderStatusCancelled
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
	ContractStatusRenewed
)


//==============================================================
// CaseStatus Declaration
//==============================================================
type CaseStatus int
const (
    CaseStatusNew CaseStatus = iota
	CaseStatusOpen
	CaseStatusPendingCustomer
	CaseStatusPendingExternal
	CaseStatusOnHold
	CaseStatusResolved
	CaseStatusClosed
	CaseStatusReopened
)


//==============================================================
// CasePriority Declaration
//==============================================================
type CasePriority int
const (
    CasePriorityLow CasePriority = iota
	CasePriorityMedium
	CasePriorityHigh
	CasePriorityCritical
)


//==============================================================
// CaseOrigin Declaration
//==============================================================
type CaseOrigin int
const (
    CaseOriginEmail CaseOrigin = iota
	CaseOriginPhone
	CaseOriginWeb
	CaseOriginChat
	CaseOriginSocial
	CaseOriginCommunity
)


//==============================================================
// CaseSeverity Declaration
//==============================================================
type CaseSeverity int
const (
    CaseSeverityMinor CaseSeverity = iota
	CaseSeverityMajor
	CaseSeverityCritical
	CaseSeverityBlocker
)


//==============================================================
// ActivityType Declaration
//==============================================================
type ActivityType int
const (
    ActivityTypeTask ActivityType = iota
	ActivityTypeCall
	ActivityTypeMeeting
	ActivityTypeDemo
	ActivityTypeFollowUp
)


//==============================================================
// ActivityStatus Declaration
//==============================================================
type ActivityStatus int
const (
    ActivityStatusNotStarted ActivityStatus = iota
	ActivityStatusInProgress
	ActivityStatusCompleted
	ActivityStatusDeferred
	ActivityStatusCancelled
)


//==============================================================
// ActivityPriority Declaration
//==============================================================
type ActivityPriority int
const (
    ActivityPriorityLow ActivityPriority = iota
	ActivityPriorityNormal
	ActivityPriorityHigh
	ActivityPriorityUrgent
)


//==============================================================
// CampaignStatus Declaration
//==============================================================
type CampaignStatus int
const (
    CampaignStatusPlanned CampaignStatus = iota
	CampaignStatusInProgress
	CampaignStatusCompleted
	CampaignStatusOnHold
	CampaignStatusCancelled
)


//==============================================================
// CampaignType Declaration
//==============================================================
type CampaignType int
const (
    CampaignTypeEmail CampaignType = iota
	CampaignTypeSocial
	CampaignTypeEvent
	CampaignTypeWebinar
	CampaignTypeAdvertising
	CampaignTypeContentMarketing
	CampaignTypeReferral
)


//==============================================================
// CampaignMemberStatus Declaration
//==============================================================
type CampaignMemberStatus int
const (
    CampaignMemberStatusSent CampaignMemberStatus = iota
	CampaignMemberStatusOpened
	CampaignMemberStatusResponded
	CampaignMemberStatusUnsubscribed
	CampaignMemberStatusBounced
	CampaignMemberStatusRegistered
	CampaignMemberStatusAttended
	CampaignMemberStatusNoShow
)


//==============================================================
// CampaignMemberType Declaration
//==============================================================
type CampaignMemberType int
const (
    CampaignMemberTypeLead CampaignMemberType = iota
	CampaignMemberTypeContact
)


//==============================================================
// EmailDirection Declaration
//==============================================================
type EmailDirection int
const (
    EmailDirectionInbound EmailDirection = iota
	EmailDirectionOutbound
	EmailDirectionInternal
)


//==============================================================
// EmailStatus Declaration
//==============================================================
type EmailStatus int
const (
    EmailStatusDraft EmailStatus = iota
	EmailStatusSent
	EmailStatusDelivered
	EmailStatusOpened
	EmailStatusBounced
	EmailStatusFailed
	EmailStatusReplied
)

