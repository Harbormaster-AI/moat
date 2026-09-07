package model


//==============================================================
// CampaignStatus Declaration
//==============================================================
type CampaignStatus int
const (
    CampaignStatusDraft CampaignStatus = iota
	CampaignStatusActive
	CampaignStatusPaused
	CampaignStatusCompleted
	CampaignStatusCancelled
)


//==============================================================
// LineItemStatus Declaration
//==============================================================
type LineItemStatus int
const (
    LineItemStatusDraft LineItemStatus = iota
	LineItemStatusScheduled
	LineItemStatusRunning
	LineItemStatusPaused
	LineItemStatusCompleted
	LineItemStatusCancelled
)


//==============================================================
// ObjectiveType Declaration
//==============================================================
type ObjectiveType int
const (
    ObjectiveTypeAwareness ObjectiveType = iota
	ObjectiveTypeReach
	ObjectiveTypeTraffic
	ObjectiveTypeEngagement
	ObjectiveTypeLeads
	ObjectiveTypeSales
	ObjectiveTypeAppInstalls
	ObjectiveTypeVideoViews
)


//==============================================================
// PricingModel Declaration
//==============================================================
type PricingModel int
const (
    PricingModelCPM PricingModel = iota
	PricingModelCPC
	PricingModelCPA
	PricingModelCPL
	PricingModelCPV
	PricingModelFlatFee
)


//==============================================================
// BidStrategyType Declaration
//==============================================================
type BidStrategyType int
const (
    BidStrategyTypeManual BidStrategyType = iota
	BidStrategyTypeAutoMaximizeClicks
	BidStrategyTypeAutoTargetCPA
	BidStrategyTypeAutoTargetROAS
)


//==============================================================
// PacingType Declaration
//==============================================================
type PacingType int
const (
    PacingTypeEven PacingType = iota
	PacingTypeASAP
	PacingTypeSmooth
)


//==============================================================
// FrequencyScope Declaration
//==============================================================
type FrequencyScope int
const (
    FrequencyScopeCampaign FrequencyScope = iota
	FrequencyScopeLineItem
	FrequencyScopeCreative
)


//==============================================================
// FrequencyPeriod Declaration
//==============================================================
type FrequencyPeriod int
const (
    FrequencyPeriodHour FrequencyPeriod = iota
	FrequencyPeriodDay
	FrequencyPeriodWeek
	FrequencyPeriodMonth
	FrequencyPeriodLifetime
)


//==============================================================
// ChannelType Declaration
//==============================================================
type ChannelType int
const (
    ChannelTypeProgrammatic ChannelType = iota
	ChannelTypeDirect
	ChannelTypeSearch
	ChannelTypeSocial
	ChannelTypeEmail
	ChannelTypeAffiliate
	ChannelTypeDOOH
)


//==============================================================
// AdFormat Declaration
//==============================================================
type AdFormat int
const (
    AdFormatBanner AdFormat = iota
	AdFormatVideo
	AdFormatNative
	AdFormatAudio
	AdFormatInterstitial
	AdFormatRichMedia
	AdFormatSearchText
	AdFormatSocialPost
	AdFormatCTVVideo
)


//==============================================================
// CreativeType Declaration
//==============================================================
type CreativeType int
const (
    CreativeTypeImage CreativeType = iota
	CreativeTypeVideo
	CreativeTypeHTML5
	CreativeTypeAudio
)


//==============================================================
// DeviceType Declaration
//==============================================================
type DeviceType int
const (
    DeviceTypeDesktop DeviceType = iota
	DeviceTypeMobile
	DeviceTypeTablet
	DeviceTypeConnectedTV
)


//==============================================================
// PlatformType Declaration
//==============================================================
type PlatformType int
const (
    PlatformTypeWeb PlatformType = iota
	PlatformTypeMobileApp
	PlatformTypeCTV
)


//==============================================================
// TargetingOperator Declaration
//==============================================================
type TargetingOperator int
const (
    TargetingOperatorInclude TargetingOperator = iota
	TargetingOperatorExclude
)


//==============================================================
// DealType Declaration
//==============================================================
type DealType int
const (
    DealTypeOpenAuction DealType = iota
	DealTypePrivateAuction
	DealTypePreferredDeal
	DealTypeProgrammaticGuaranteed
)


//==============================================================
// PublisherType Declaration
//==============================================================
type PublisherType int
const (
    PublisherTypeSite PublisherType = iota
	PublisherTypeApp
	PublisherTypeNetwork
	PublisherTypeCTVApp
)


//==============================================================
// DataProviderType Declaration
//==============================================================
type DataProviderType int
const (
    DataProviderTypeFirstParty DataProviderType = iota
	DataProviderTypeSecondParty
	DataProviderTypeThirdParty
)


//==============================================================
// BrandSafetyLevel Declaration
//==============================================================
type BrandSafetyLevel int
const (
    BrandSafetyLevelNone BrandSafetyLevel = iota
	BrandSafetyLevelModerate
	BrandSafetyLevelStrict
)


//==============================================================
// ContentRating Declaration
//==============================================================
type ContentRating int
const (
    ContentRatingG ContentRating = iota
	ContentRatingPG
	ContentRatingPGThirteen
	ContentRatingR
	ContentRatingMature
	ContentRatingUnrated
)


//==============================================================
// MetricType Declaration
//==============================================================
type MetricType int
const (
    MetricTypeImpressions MetricType = iota
	MetricTypeViewableImpressions
	MetricTypeClicks
	MetricTypeCTR
	MetricTypeReach
	MetricTypeFrequency
	MetricTypeVideoStarts
	MetricTypeVideoCompletions
	MetricTypeAvgViewTime
	MetricTypeConversions
	MetricTypeViewThroughConversions
	MetricTypeSpend
	MetricTypeCPM
	MetricTypeCPC
	MetricTypeCPA
)


//==============================================================
// AttributionModel Declaration
//==============================================================
type AttributionModel int
const (
    AttributionModelLastClick AttributionModel = iota
	AttributionModelFirstTouch
	AttributionModelLinear
	AttributionModelTimeDecay
	AttributionModelPositionBased
	AttributionModelDataDriven
)


//==============================================================
// IOStatus Declaration
//==============================================================
type IOStatus int
const (
    IOStatusDraft IOStatus = iota
	IOStatusSent
	IOStatusExecuted
	IOStatusOnHold
	IOStatusClosed
	IOStatusCancelled
)


//==============================================================
// ReportType Declaration
//==============================================================
type ReportType int
const (
    ReportTypePerformance ReportType = iota
	ReportTypeDelivery
	ReportTypeInventory
	ReportTypeBilling
)


//==============================================================
// PixelType Declaration
//==============================================================
type PixelType int
const (
    PixelTypeImage PixelType = iota
	PixelTypeJavaScript
	PixelTypeServerSide
)


//==============================================================
// ConversionEventType Declaration
//==============================================================
type ConversionEventType int
const (
    ConversionEventTypeLead ConversionEventType = iota
	ConversionEventTypePurchase
	ConversionEventTypeSignup
	ConversionEventTypeAddToCart
	ConversionEventTypeViewContent
	ConversionEventTypeAppInstall
)


//==============================================================
// GeoRegionType Declaration
//==============================================================
type GeoRegionType int
const (
    GeoRegionTypeCountry GeoRegionType = iota
	GeoRegionTypeState
	GeoRegionTypeProvince
	GeoRegionTypeCity
	GeoRegionTypeDMA
	GeoRegionTypePostalCode
)


//==============================================================
// AccountRole Declaration
//==============================================================
type AccountRole int
const (
    AccountRoleAdmin AccountRole = iota
	AccountRoleTrader
	AccountRoleAnalyst
	AccountRoleViewer
)


//==============================================================
// PaymentTerms Declaration
//==============================================================
type PaymentTerms int
const (
    PaymentTermsPrepaid PaymentTerms = iota
	PaymentTermsNetFifteen
	PaymentTermsNetThirty
	PaymentTermsNetSixty
)


//==============================================================
// PaymentMethodType Declaration
//==============================================================
type PaymentMethodType int
const (
    PaymentMethodTypeCreditCard PaymentMethodType = iota
	PaymentMethodTypeInvoice
	PaymentMethodTypeWire
	PaymentMethodTypeACH
)


//==============================================================
// InventoryType Declaration
//==============================================================
type InventoryType int
const (
    InventoryTypeDisplay InventoryType = iota
	InventoryTypeVideo
	InventoryTypeNative
	InventoryTypeAudio
	InventoryTypeSearch
	InventoryTypeSocial
	InventoryTypeDOOH
)


//==============================================================
// CreativeApprovalStatus Declaration
//==============================================================
type CreativeApprovalStatus int
const (
    CreativeApprovalStatusPending CreativeApprovalStatus = iota
	CreativeApprovalStatusApproved
	CreativeApprovalStatusRejected
)


//==============================================================
// ExperimentStatus Declaration
//==============================================================
type ExperimentStatus int
const (
    ExperimentStatusPlanned ExperimentStatus = iota
	ExperimentStatusRunning
	ExperimentStatusPaused
	ExperimentStatusCompleted
	ExperimentStatusCancelled
)

