package model


//==============================================================
// AircraftType Declaration
//==============================================================
type AircraftType int
const (
    AircraftTypeNarrowBody AircraftType = iota
	AircraftTypeWideBody
	AircraftTypeRegionalJet
	AircraftTypeTurboprop
	AircraftTypeBusinessJet
	AircraftTypeHelicopter
	AircraftTypeeVTOL
	AircraftTypeCargoPlane
)


//==============================================================
// EngineCategory Declaration
//==============================================================
type EngineCategory int
const (
    EngineCategoryTurbofan EngineCategory = iota
	EngineCategoryTurboprop
	EngineCategoryTurbojet
	EngineCategoryPiston
	EngineCategoryElectric
	EngineCategoryRocket
)


//==============================================================
// LandingGearType Declaration
//==============================================================
type LandingGearType int
const (
    LandingGearTypeTricycle LandingGearType = iota
	LandingGearTypeTandem
	LandingGearTypeTaildragger
	LandingGearTypeSkid
	LandingGearTypeFloats
	LandingGearTypeRetractable
)


//==============================================================
// OptionCategory Declaration
//==============================================================
type OptionCategory int
const (
    OptionCategoryCabin OptionCategory = iota
	OptionCategoryConnectivity
	OptionCategorySafety
	OptionCategoryPerformance
	OptionCategoryPaint
	OptionCategoryFlightDeck
)


//==============================================================
// PackageType Declaration
//==============================================================
type PackageType int
const (
    PackageTypePerformancePack PackageType = iota
	PackageTypeCabinPack
	PackageTypeConnectivityPack
	PackageTypeCompliancePack
)


//==============================================================
// SupplierType Declaration
//==============================================================
type SupplierType int
const (
    SupplierTypeAirframe SupplierType = iota
	SupplierTypeEngine
	SupplierTypeAvionics
	SupplierTypeSystems
	SupplierTypeMaterials
	SupplierTypeMRO
	SupplierTypeTesting
)


//==============================================================
// SupplierApprovalStatus Declaration
//==============================================================
type SupplierApprovalStatus int
const (
    SupplierApprovalStatusApplied SupplierApprovalStatus = iota
	SupplierApprovalStatusApproved
	SupplierApprovalStatusOnHold
	SupplierApprovalStatusSuspended
)


//==============================================================
// ComponentCategory Declaration
//==============================================================
type ComponentCategory int
const (
    ComponentCategoryStructure ComponentCategory = iota
	ComponentCategorySystem
	ComponentCategoryAvionics
	ComponentCategoryInterior
	ComponentCategoryLandingGear
	ComponentCategoryPowerplant
	ComponentCategoryConsumable
)


//==============================================================
// SerializationMethod Declaration
//==============================================================
type SerializationMethod int
const (
    SerializationMethodSerialized SerializationMethod = iota
	SerializationMethodLotTracked
	SerializationMethodNone
)


//==============================================================
// ProductionLineType Declaration
//==============================================================
type ProductionLineType int
const (
    ProductionLineTypeFinalAssembly ProductionLineType = iota
	ProductionLineTypeSubAssembly
	ProductionLineTypeIntegration
	ProductionLineTypeTestAndDelivery
)


//==============================================================
// ProductionOrderStatus Declaration
//==============================================================
type ProductionOrderStatus int
const (
    ProductionOrderStatusPlanned ProductionOrderStatus = iota
	ProductionOrderStatusReleased
	ProductionOrderStatusInAssembly
	ProductionOrderStatusFlightTest
	ProductionOrderStatusCompleted
)


//==============================================================
// ScheduleStatus Declaration
//==============================================================
type ScheduleStatus int
const (
    ScheduleStatusDraft ScheduleStatus = iota
	ScheduleStatusPublished
	ScheduleStatusRevised
	ScheduleStatusClosed
)


//==============================================================
// OperatorType Declaration
//==============================================================
type OperatorType int
const (
    OperatorTypeAirline OperatorType = iota
	OperatorTypeCargo
	OperatorTypeGovernment
	OperatorTypePrivate
	OperatorTypeLessor
)


//==============================================================
// AircraftOrderStatus Declaration
//==============================================================
type AircraftOrderStatus int
const (
    AircraftOrderStatusDraft AircraftOrderStatus = iota
	AircraftOrderStatusCommitted
	AircraftOrderStatusInProduction
	AircraftOrderStatusDelivered
	AircraftOrderStatusCancelled
)


//==============================================================
// WarrantyType Declaration
//==============================================================
type WarrantyType int
const (
    WarrantyTypeBasic WarrantyType = iota
	WarrantyTypePowerplant
	WarrantyTypeAvionics
	WarrantyTypeCorrosion
)


//==============================================================
// AppointmentStatus Declaration
//==============================================================
type AppointmentStatus int
const (
    AppointmentStatusScheduled AppointmentStatus = iota
	AppointmentStatusInProgress
	AppointmentStatusCompleted
	AppointmentStatusCancelled
	AppointmentStatusDeferred
)


//==============================================================
// WorkOrderStatus Declaration
//==============================================================
type WorkOrderStatus int
const (
    WorkOrderStatusOpen WorkOrderStatus = iota
	WorkOrderStatusInProgress
	WorkOrderStatusAwaitingParts
	WorkOrderStatusClosed
	WorkOrderStatusDeferred
)


//==============================================================
// ServiceBulletinCategory Declaration
//==============================================================
type ServiceBulletinCategory int
const (
    ServiceBulletinCategoryRecommended ServiceBulletinCategory = iota
	ServiceBulletinCategoryOptional
	ServiceBulletinCategoryAlert
	ServiceBulletinCategoryMandatory
)


//==============================================================
// ConnectivityStatus Declaration
//==============================================================
type ConnectivityStatus int
const (
    ConnectivityStatusOffline ConnectivityStatus = iota
	ConnectivityStatusOnline
	ConnectivityStatusDegraded
)


//==============================================================
// EventSeverity Declaration
//==============================================================
type EventSeverity int
const (
    EventSeverityInfo EventSeverity = iota
	EventSeverityWarning
	EventSeverityCritical
)


//==============================================================
// SoftwareLoadType Declaration
//==============================================================
type SoftwareLoadType int
const (
    SoftwareLoadTypeFlightDeckSoftware SoftwareLoadType = iota
	SoftwareLoadTypeMaintenanceTools
	SoftwareLoadTypeCabinIFE
	SoftwareLoadTypeConnectivityModem
)


//==============================================================
// SalesCampaignStatus Declaration
//==============================================================
type SalesCampaignStatus int
const (
    SalesCampaignStatusProspecting SalesCampaignStatus = iota
	SalesCampaignStatusProposal
	SalesCampaignStatusNegotiation
	SalesCampaignStatusWon
	SalesCampaignStatusLost
)


//==============================================================
// ProgramStatus Declaration
//==============================================================
type ProgramStatus int
const (
    ProgramStatusConcept ProgramStatus = iota
	ProgramStatusDevelopment
	ProgramStatusCertification
	ProgramStatusProduction
	ProgramStatusInService
	ProgramStatusSunset
)

