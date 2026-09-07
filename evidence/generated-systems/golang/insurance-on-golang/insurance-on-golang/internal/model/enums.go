package model


//==============================================================
// LineOfBusiness Declaration
//==============================================================
type LineOfBusiness int
const (
    LineOfBusinessPersonalAuto LineOfBusiness = iota
	LineOfBusinessHomeowners
	LineOfBusinessRenters
	LineOfBusinessTermLife
	LineOfBusinessWholeLife
	LineOfBusinessCommercialProperty
	LineOfBusinessGeneralLiability
	LineOfBusinessWorkersCompensation
)


//==============================================================
// CoverageType Declaration
//==============================================================
type CoverageType int
const (
    CoverageTypeLiability CoverageType = iota
	CoverageTypeCollision
	CoverageTypeComprehensive
	CoverageTypePropertyDamage
	CoverageTypeBodilyInjury
	CoverageTypeUninsuredMotorist
	CoverageTypeMedicalPayments
	CoverageTypeDwelling
	CoverageTypeContents
	CoverageTypePersonalLiability
	CoverageTypeBusinessInterruption
	CoverageTypeProfessionalLiability
)


//==============================================================
// DistributionChannelType Declaration
//==============================================================
type DistributionChannelType int
const (
    DistributionChannelTypeAgency DistributionChannelType = iota
	DistributionChannelTypeBroker
	DistributionChannelTypeDirect
	DistributionChannelTypeBancassurance
	DistributionChannelTypeAffinityPartner
	DistributionChannelTypeOnlineAggregator
)


//==============================================================
// ProducerStatus Declaration
//==============================================================
type ProducerStatus int
const (
    ProducerStatusActive ProducerStatus = iota
	ProducerStatusSuspended
	ProducerStatusTerminated
)


//==============================================================
// CustomerType Declaration
//==============================================================
type CustomerType int
const (
    CustomerTypeIndividual CustomerType = iota
	CustomerTypeOrganization
)


//==============================================================
// ApplicationStatus Declaration
//==============================================================
type ApplicationStatus int
const (
    ApplicationStatusDraft ApplicationStatus = iota
	ApplicationStatusSubmitted
	ApplicationStatusUnderReview
	ApplicationStatusQuoted
	ApplicationStatusDeclined
	ApplicationStatusWithdrawn
	ApplicationStatusBound
)


//==============================================================
// UnderwritingDecisionType Declaration
//==============================================================
type UnderwritingDecisionType int
const (
    UnderwritingDecisionTypeApprove UnderwritingDecisionType = iota
	UnderwritingDecisionTypeConditionalApprove
	UnderwritingDecisionTypeRefer
	UnderwritingDecisionTypeDecline
)


//==============================================================
// PolicyStatus Declaration
//==============================================================
type PolicyStatus int
const (
    PolicyStatusQuoted PolicyStatus = iota
	PolicyStatusActive
	PolicyStatusLapsed
	PolicyStatusCancelled
	PolicyStatusExpired
	PolicyStatusPendingCancel
	PolicyStatusPendingReinstatement
)


//==============================================================
// PaymentPlanType Declaration
//==============================================================
type PaymentPlanType int
const (
    PaymentPlanTypeAnnual PaymentPlanType = iota
	PaymentPlanTypeSemiAnnual
	PaymentPlanTypeQuarterly
	PaymentPlanTypeMonthly
	PaymentPlanTypePayInFull
)


//==============================================================
// BillingStatus Declaration
//==============================================================
type BillingStatus int
const (
    BillingStatusCurrent BillingStatus = iota
	BillingStatusDelinquent
	BillingStatusCollections
	BillingStatusClosed
)


//==============================================================
// InvoiceStatus Declaration
//==============================================================
type InvoiceStatus int
const (
    InvoiceStatusOpen InvoiceStatus = iota
	InvoiceStatusPaid
	InvoiceStatusPartiallyPaid
	InvoiceStatusVoid
)


//==============================================================
// PaymentMethod Declaration
//==============================================================
type PaymentMethod int
const (
    PaymentMethodACH PaymentMethod = iota
	PaymentMethodCreditCard
	PaymentMethodDebitCard
	PaymentMethodCheck
	PaymentMethodCash
	PaymentMethodWire
)


//==============================================================
// PaymentStatus Declaration
//==============================================================
type PaymentStatus int
const (
    PaymentStatusPending PaymentStatus = iota
	PaymentStatusSettled
	PaymentStatusFailed
	PaymentStatusRefunded
	PaymentStatusReversed
)


//==============================================================
// InsuredObjectType Declaration
//==============================================================
type InsuredObjectType int
const (
    InsuredObjectTypeVehicle InsuredObjectType = iota
	InsuredObjectTypeProperty
	InsuredObjectTypePerson
	InsuredObjectTypeEquipment
	InsuredObjectTypeLiabilityExposure
)


//==============================================================
// RelationshipType Declaration
//==============================================================
type RelationshipType int
const (
    RelationshipTypeSpouse RelationshipType = iota
	RelationshipTypeChild
	RelationshipTypeParent
	RelationshipTypeSibling
	RelationshipTypeBusinessPartner
	RelationshipTypeEstate
	RelationshipTypeTrust
	RelationshipTypeOther
)


//==============================================================
// ClaimStatus Declaration
//==============================================================
type ClaimStatus int
const (
    ClaimStatusOpen ClaimStatus = iota
	ClaimStatusClosed
	ClaimStatusReopened
	ClaimStatusDenied
	ClaimStatusPendingInvestigation
	ClaimStatusLitigation
)


//==============================================================
// PerilType Declaration
//==============================================================
type PerilType int
const (
    PerilTypeAutoAccident PerilType = iota
	PerilTypeFire
	PerilTypeTheft
	PerilTypeWindstorm
	PerilTypeFlood
	PerilTypeHail
	PerilTypeEarthquake
	PerilTypeVandalism
	PerilTypeInjury
	PerilTypeDeath
)


//==============================================================
// CauseOfLoss Declaration
//==============================================================
type CauseOfLoss int
const (
    CauseOfLossCollision CauseOfLoss = iota
	CauseOfLossWeather
	CauseOfLossMechanicalFailure
	CauseOfLossHumanError
	CauseOfLossNaturalDisaster
	CauseOfLossTheft
	CauseOfLossVandalism
	CauseOfLossLiabilityClaim
	CauseOfLossIllness
)


//==============================================================
// ExposureType Declaration
//==============================================================
type ExposureType int
const (
    ExposureTypeBodilyInjury ExposureType = iota
	ExposureTypePropertyDamage
	ExposureTypeMedical
	ExposureTypeUninsuredMotorist
	ExposureTypePersonalInjuryProtection
	ExposureTypeDwellingDamage
	ExposureTypeContentsDamage
	ExposureTypeBusinessIncome
)


//==============================================================
// ExposureStatus Declaration
//==============================================================
type ExposureStatus int
const (
    ExposureStatusOpen ExposureStatus = iota
	ExposureStatusClosed
	ExposureStatusPending
	ExposureStatusReserved
)


//==============================================================
// AdjusterType Declaration
//==============================================================
type AdjusterType int
const (
    AdjusterTypeStaff AdjusterType = iota
	AdjusterTypeIndependent
	AdjusterTypePublic
)


//==============================================================
// ReserveType Declaration
//==============================================================
type ReserveType int
const (
    ReserveTypeIndemnity ReserveType = iota
	ReserveTypeExpense
	ReserveTypeLegal
	ReserveTypeMedical
)


//==============================================================
// ReserveStatus Declaration
//==============================================================
type ReserveStatus int
const (
    ReserveStatusOpen ReserveStatus = iota
	ReserveStatusReleased
	ReserveStatusIncreased
	ReserveStatusDecreased
	ReserveStatusClosed
)


//==============================================================
// PayeeType Declaration
//==============================================================
type PayeeType int
const (
    PayeeTypeClaimant PayeeType = iota
	PayeeTypeBeneficiary
	PayeeTypeServiceProvider
	PayeeTypeLienholder
	PayeeTypeAttorney
)


//==============================================================
// ServiceProviderType Declaration
//==============================================================
type ServiceProviderType int
const (
    ServiceProviderTypeRepairShop ServiceProviderType = iota
	ServiceProviderTypeTowing
	ServiceProviderTypeMedicalProvider
	ServiceProviderTypeAttorney
	ServiceProviderTypeForensicEngineer
	ServiceProviderTypeRentalCar
)


//==============================================================
// NetworkStatus Declaration
//==============================================================
type NetworkStatus int
const (
    NetworkStatusInNetwork NetworkStatus = iota
	NetworkStatusOutOfNetwork
)


//==============================================================
// ReinsuranceType Declaration
//==============================================================
type ReinsuranceType int
const (
    ReinsuranceTypeTreaty ReinsuranceType = iota
	ReinsuranceTypeFacultative
)


//==============================================================
// TreatyType Declaration
//==============================================================
type TreatyType int
const (
    TreatyTypeQuotaShare TreatyType = iota
	TreatyTypeSurplus
	TreatyTypeExcessOfLoss
	TreatyTypeStopLoss
)


//==============================================================
// ThirdPartyType Declaration
//==============================================================
type ThirdPartyType int
const (
    ThirdPartyTypeIndividual ThirdPartyType = iota
	ThirdPartyTypeCompany
	ThirdPartyTypeGovernmentAgency
)


//==============================================================
// DocumentType Declaration
//==============================================================
type DocumentType int
const (
    DocumentTypeApplicationForm DocumentType = iota
	DocumentTypePolicyDocument
	DocumentTypeEndorsement
	DocumentTypeInvoice
	DocumentTypeClaimForm
	DocumentTypePoliceReport
	DocumentTypeEstimate
	DocumentTypePhoto
	DocumentTypeMedicalRecord
	DocumentTypeCorrespondence
)


//==============================================================
// SubrogationStatus Declaration
//==============================================================
type SubrogationStatus int
const (
    SubrogationStatusOpen SubrogationStatus = iota
	SubrogationStatusNegotiating
	SubrogationStatusSettled
	SubrogationStatusUncollectible
	SubrogationStatusClosed
)

