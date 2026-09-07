
// enum type UserRole
export let UserRole = {
	SalesRep:"SalesRep",
	SalesManager:"SalesManager",
	ServiceAgent:"ServiceAgent",
	MarketingSpecialist:"MarketingSpecialist",
	Administrator:"Administrator",
	Executive:"Executive",
}

// enum type UserStatus
export let UserStatus = {
	Active:"Active",
	Inactive:"Inactive",
	Locked:"Locked",
	PendingInvite:"PendingInvite",
}

// enum type TeamType
export let TeamType = {
	Sales:"Sales",
	Service:"Service",
	Marketing:"Marketing",
	AccountTeam:"AccountTeam",
	DealDesk:"DealDesk",
	CrossFunctional:"CrossFunctional",
}

// enum type TerritoryType
export let TerritoryType = {
	Geographic:"Geographic",
	Industry:"Industry",
	NamedAccount:"NamedAccount",
	Segment:"Segment",
	Hybrid:"Hybrid",
}

// enum type AccountType
export let AccountType = {
	Prospect:"Prospect",
	Customer:"Customer",
	Partner:"Partner",
	Vendor:"Vendor",
	Competitor:"Competitor",
}

// enum type AccountLifecycleStage
export let AccountLifecycleStage = {
	Subscriber:"Subscriber",
	Lead:"Lead",
	MarketingQualified:"MarketingQualified",
	SalesQualified:"SalesQualified",
	Customer:"Customer",
	Evangelist:"Evangelist",
	Churned:"Churned",
}

// enum type ContactMethod
export let ContactMethod = {
	Email:"Email",
	Phone:"Phone",
	Mobile:"Mobile",
	SMS:"SMS",
	InPerson:"InPerson",
	Web:"Web",
}

// enum type LeadStatus
export let LeadStatus = {
	New:"New",
	Working:"Working",
	Nurturing:"Nurturing",
	Qualified:"Qualified",
	Unqualified:"Unqualified",
	Converted:"Converted",
}

// enum type LeadSource
export let LeadSource = {
	Web:"Web",
	Referral:"Referral",
	Event:"Event",
	Partner:"Partner",
	Advertisement:"Advertisement",
	Outbound:"Outbound",
	Inbound:"Inbound",
	Social:"Social",
	Other:"Other",
}

// enum type LeadRating
export let LeadRating = {
	Hot:"Hot",
	Warm:"Warm",
	Cold:"Cold",
}

// enum type OpportunityStage
export let OpportunityStage = {
	Qualification:"Qualification",
	Discovery:"Discovery",
	Proposal:"Proposal",
	Negotiation:"Negotiation",
	ClosedWon:"ClosedWon",
	ClosedLost:"ClosedLost",
}

// enum type OpportunityType
export let OpportunityType = {
	NewBusiness:"NewBusiness",
	ExistingBusiness:"ExistingBusiness",
	Renewal:"Renewal",
	Upsell:"Upsell",
	CrossSell:"CrossSell",
}

// enum type ForecastCategory
export let ForecastCategory = {
	Pipeline:"Pipeline",
	BestCase:"BestCase",
	Commit:"Commit",
	Omitted:"Omitted",
	Closed:"Closed",
}

// enum type ProductType
export let ProductType = {
	Good:"Good",
	Service:"Service",
	Subscription:"Subscription",
	Bundle:"Bundle",
}

// enum type UnitOfMeasure
export let UnitOfMeasure = {
	Each:"Each",
	Hour:"Hour",
	Day:"Day",
	Month:"Month",
	User:"User",
	Package:"Package",
}

// enum type QuoteStatus
export let QuoteStatus = {
	Draft:"Draft",
	Presented:"Presented",
	Approved:"Approved",
	Rejected:"Rejected",
	Accepted:"Accepted",
	Expired:"Expired",
	Withdrawn:"Withdrawn",
}

// enum type OrderStatus
export let OrderStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	PartiallyFulfilled:"PartiallyFulfilled",
	Fulfilled:"Fulfilled",
	Invoiced:"Invoiced",
	Cancelled:"Cancelled",
}

// enum type ContractStatus
export let ContractStatus = {
	Draft:"Draft",
	Active:"Active",
	Suspended:"Suspended",
	Expired:"Expired",
	Terminated:"Terminated",
	Renewed:"Renewed",
}

// enum type CaseStatus
export let CaseStatus = {
	New:"New",
	Open:"Open",
	PendingCustomer:"PendingCustomer",
	PendingExternal:"PendingExternal",
	OnHold:"OnHold",
	Resolved:"Resolved",
	Closed:"Closed",
	Reopened:"Reopened",
}

// enum type CasePriority
export let CasePriority = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Critical:"Critical",
}

// enum type CaseOrigin
export let CaseOrigin = {
	Email:"Email",
	Phone:"Phone",
	Web:"Web",
	Chat:"Chat",
	Social:"Social",
	Community:"Community",
}

// enum type CaseSeverity
export let CaseSeverity = {
	Minor:"Minor",
	Major:"Major",
	Critical:"Critical",
	Blocker:"Blocker",
}

// enum type ActivityType
export let ActivityType = {
	Task:"Task",
	Call:"Call",
	Meeting:"Meeting",
	Demo:"Demo",
	FollowUp:"FollowUp",
}

// enum type ActivityStatus
export let ActivityStatus = {
	NotStarted:"NotStarted",
	InProgress:"InProgress",
	Completed:"Completed",
	Deferred:"Deferred",
	Cancelled:"Cancelled",
}

// enum type ActivityPriority
export let ActivityPriority = {
	Low:"Low",
	Normal:"Normal",
	High:"High",
	Urgent:"Urgent",
}

// enum type CampaignStatus
export let CampaignStatus = {
	Planned:"Planned",
	InProgress:"InProgress",
	Completed:"Completed",
	OnHold:"OnHold",
	Cancelled:"Cancelled",
}

// enum type CampaignType
export let CampaignType = {
	Email:"Email",
	Social:"Social",
	Event:"Event",
	Webinar:"Webinar",
	Advertising:"Advertising",
	ContentMarketing:"ContentMarketing",
	Referral:"Referral",
}

// enum type CampaignMemberStatus
export let CampaignMemberStatus = {
	Sent:"Sent",
	Opened:"Opened",
	Responded:"Responded",
	Unsubscribed:"Unsubscribed",
	Bounced:"Bounced",
	Registered:"Registered",
	Attended:"Attended",
	NoShow:"NoShow",
}

// enum type CampaignMemberType
export let CampaignMemberType = {
	Lead:"Lead",
	Contact:"Contact",
}

// enum type EmailDirection
export let EmailDirection = {
	Inbound:"Inbound",
	Outbound:"Outbound",
	Internal:"Internal",
}

// enum type EmailStatus
export let EmailStatus = {
	Draft:"Draft",
	Sent:"Sent",
	Delivered:"Delivered",
	Opened:"Opened",
	Bounced:"Bounced",
	Failed:"Failed",
	Replied:"Replied",
}
