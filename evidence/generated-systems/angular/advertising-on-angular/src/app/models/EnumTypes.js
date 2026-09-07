
// enum type CampaignStatus
export let CampaignStatus = {
	Draft:"Draft",
	Active:"Active",
	Paused:"Paused",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type LineItemStatus
export let LineItemStatus = {
	Draft:"Draft",
	Scheduled:"Scheduled",
	Running:"Running",
	Paused:"Paused",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type ObjectiveType
export let ObjectiveType = {
	Awareness:"Awareness",
	Reach:"Reach",
	Traffic:"Traffic",
	Engagement:"Engagement",
	Leads:"Leads",
	Sales:"Sales",
	AppInstalls:"AppInstalls",
	VideoViews:"VideoViews",
}

// enum type PricingModel
export let PricingModel = {
	CPM:"CPM",
	CPC:"CPC",
	CPA:"CPA",
	CPL:"CPL",
	CPV:"CPV",
	FlatFee:"FlatFee",
}

// enum type BidStrategyType
export let BidStrategyType = {
	Manual:"Manual",
	AutoMaximizeClicks:"AutoMaximizeClicks",
	AutoTargetCPA:"AutoTargetCPA",
	AutoTargetROAS:"AutoTargetROAS",
}

// enum type PacingType
export let PacingType = {
	Even:"Even",
	ASAP:"ASAP",
	Smooth:"Smooth",
}

// enum type FrequencyScope
export let FrequencyScope = {
	Campaign:"Campaign",
	LineItem:"LineItem",
	Creative:"Creative",
}

// enum type FrequencyPeriod
export let FrequencyPeriod = {
	Hour:"Hour",
	Day:"Day",
	Week:"Week",
	Month:"Month",
	Lifetime:"Lifetime",
}

// enum type ChannelType
export let ChannelType = {
	Programmatic:"Programmatic",
	Direct:"Direct",
	Search:"Search",
	Social:"Social",
	Email:"Email",
	Affiliate:"Affiliate",
	DOOH:"DOOH",
}

// enum type AdFormat
export let AdFormat = {
	Banner:"Banner",
	Video:"Video",
	Native:"Native",
	Audio:"Audio",
	Interstitial:"Interstitial",
	RichMedia:"RichMedia",
	SearchText:"SearchText",
	SocialPost:"SocialPost",
	CTVVideo:"CTVVideo",
}

// enum type CreativeType
export let CreativeType = {
	Image:"Image",
	Video:"Video",
	HTML5:"HTML5",
	Audio:"Audio",
}

// enum type DeviceType
export let DeviceType = {
	Desktop:"Desktop",
	Mobile:"Mobile",
	Tablet:"Tablet",
	ConnectedTV:"ConnectedTV",
}

// enum type PlatformType
export let PlatformType = {
	Web:"Web",
	MobileApp:"MobileApp",
	CTV:"CTV",
}

// enum type TargetingOperator
export let TargetingOperator = {
	Include:"Include",
	Exclude:"Exclude",
}

// enum type DealType
export let DealType = {
	OpenAuction:"OpenAuction",
	PrivateAuction:"PrivateAuction",
	PreferredDeal:"PreferredDeal",
	ProgrammaticGuaranteed:"ProgrammaticGuaranteed",
}

// enum type PublisherType
export let PublisherType = {
	Site:"Site",
	App:"App",
	Network:"Network",
	CTVApp:"CTVApp",
}

// enum type DataProviderType
export let DataProviderType = {
	FirstParty:"FirstParty",
	SecondParty:"SecondParty",
	ThirdParty:"ThirdParty",
}

// enum type BrandSafetyLevel
export let BrandSafetyLevel = {
	None:"None",
	Moderate:"Moderate",
	Strict:"Strict",
}

// enum type ContentRating
export let ContentRating = {
	G:"G",
	PG:"PG",
	PGThirteen:"PGThirteen",
	R:"R",
	Mature:"Mature",
	Unrated:"Unrated",
}

// enum type MetricType
export let MetricType = {
	Impressions:"Impressions",
	ViewableImpressions:"ViewableImpressions",
	Clicks:"Clicks",
	CTR:"CTR",
	Reach:"Reach",
	Frequency:"Frequency",
	VideoStarts:"VideoStarts",
	VideoCompletions:"VideoCompletions",
	AvgViewTime:"AvgViewTime",
	Conversions:"Conversions",
	ViewThroughConversions:"ViewThroughConversions",
	Spend:"Spend",
	CPM:"CPM",
	CPC:"CPC",
	CPA:"CPA",
}

// enum type AttributionModel
export let AttributionModel = {
	LastClick:"LastClick",
	FirstTouch:"FirstTouch",
	Linear:"Linear",
	TimeDecay:"TimeDecay",
	PositionBased:"PositionBased",
	DataDriven:"DataDriven",
}

// enum type IOStatus
export let IOStatus = {
	Draft:"Draft",
	Sent:"Sent",
	Executed:"Executed",
	OnHold:"OnHold",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type ReportType
export let ReportType = {
	Performance:"Performance",
	Delivery:"Delivery",
	Inventory:"Inventory",
	Billing:"Billing",
}

// enum type PixelType
export let PixelType = {
	Image:"Image",
	JavaScript:"JavaScript",
	ServerSide:"ServerSide",
}

// enum type ConversionEventType
export let ConversionEventType = {
	Lead:"Lead",
	Purchase:"Purchase",
	Signup:"Signup",
	AddToCart:"AddToCart",
	ViewContent:"ViewContent",
	AppInstall:"AppInstall",
}

// enum type GeoRegionType
export let GeoRegionType = {
	Country:"Country",
	State:"State",
	Province:"Province",
	City:"City",
	DMA:"DMA",
	PostalCode:"PostalCode",
}

// enum type AccountRole
export let AccountRole = {
	Admin:"Admin",
	Trader:"Trader",
	Analyst:"Analyst",
	Viewer:"Viewer",
}

// enum type PaymentTerms
export let PaymentTerms = {
	Prepaid:"Prepaid",
	NetFifteen:"NetFifteen",
	NetThirty:"NetThirty",
	NetSixty:"NetSixty",
}

// enum type PaymentMethodType
export let PaymentMethodType = {
	CreditCard:"CreditCard",
	Invoice:"Invoice",
	Wire:"Wire",
	ACH:"ACH",
}

// enum type InventoryType
export let InventoryType = {
	Display:"Display",
	Video:"Video",
	Native:"Native",
	Audio:"Audio",
	Search:"Search",
	Social:"Social",
	DOOH:"DOOH",
}

// enum type CreativeApprovalStatus
export let CreativeApprovalStatus = {
	Pending:"Pending",
	Approved:"Approved",
	Rejected:"Rejected",
}

// enum type ExperimentStatus
export let ExperimentStatus = {
	Planned:"Planned",
	Running:"Running",
	Paused:"Paused",
	Completed:"Completed",
	Cancelled:"Cancelled",
}
