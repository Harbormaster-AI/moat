package model


//==============================================================
// BusinessUnitCategory Declaration
//==============================================================
type BusinessUnitCategory int
const (
    BusinessUnitCategoryConsumerGoods BusinessUnitCategory = iota
	BusinessUnitCategoryIndustrialEquipment
	BusinessUnitCategoryElectronics
	BusinessUnitCategoryPharmaceuticals
	BusinessUnitCategoryFoodBeverage
)


//==============================================================
// ProductionLineType Declaration
//==============================================================
type ProductionLineType int
const (
    ProductionLineTypeDiscrete ProductionLineType = iota
	ProductionLineTypeBatch
	ProductionLineTypeContinuous
	ProductionLineTypeFlexibleCell
)


//==============================================================
// WorkCenterType Declaration
//==============================================================
type WorkCenterType int
const (
    WorkCenterTypeMachining WorkCenterType = iota
	WorkCenterTypeAssembly
	WorkCenterTypePainting
	WorkCenterTypePackaging
	WorkCenterTypeTest
	WorkCenterTypeWarehouse
)


//==============================================================
// ItemType Declaration
//==============================================================
type ItemType int
const (
    ItemTypeFinishedGood ItemType = iota
	ItemTypeSubassembly
	ItemTypeComponent
	ItemTypeRawMaterial
	ItemTypeConsumable
	ItemTypeService
)


//==============================================================
// ProcurementType Declaration
//==============================================================
type ProcurementType int
const (
    ProcurementTypeMakeToStock ProcurementType = iota
	ProcurementTypeMakeToOrder
	ProcurementTypePurchase
	ProcurementTypeKanban
	ProcurementTypeOutsourced
)


//==============================================================
// UnitOfMeasure Declaration
//==============================================================
type UnitOfMeasure int
const (
    UnitOfMeasureEach UnitOfMeasure = iota
	UnitOfMeasureKilogram
	UnitOfMeasureGram
	UnitOfMeasurePound
	UnitOfMeasureLiter
	UnitOfMeasureMeter
	UnitOfMeasureCentimeter
	UnitOfMeasureMillimeter
	UnitOfMeasureHour
	UnitOfMeasureMinute
	UnitOfMeasureBox
	UnitOfMeasurePallet
)


//==============================================================
// TimeUnit Declaration
//==============================================================
type TimeUnit int
const (
    TimeUnitSecond TimeUnit = iota
	TimeUnitMinute
	TimeUnitHour
	TimeUnitDay
)


//==============================================================
// ProductLifecycleStatus Declaration
//==============================================================
type ProductLifecycleStatus int
const (
    ProductLifecycleStatusActive ProductLifecycleStatus = iota
	ProductLifecycleStatusPendingApproval
	ProductLifecycleStatusDiscontinued
	ProductLifecycleStatusObsolete
)


//==============================================================
// BOMStatus Declaration
//==============================================================
type BOMStatus int
const (
    BOMStatusDraft BOMStatus = iota
	BOMStatusReleased
	BOMStatusObsolete
)


//==============================================================
// RoutingType Declaration
//==============================================================
type RoutingType int
const (
    RoutingTypeStandard RoutingType = iota
	RoutingTypeAlternate
	RoutingTypeRework
)


//==============================================================
// RoutingStatus Declaration
//==============================================================
type RoutingStatus int
const (
    RoutingStatusDraft RoutingStatus = iota
	RoutingStatusReleased
	RoutingStatusObsolete
)


//==============================================================
// OperationType Declaration
//==============================================================
type OperationType int
const (
    OperationTypeSetup OperationType = iota
	OperationTypeRun
	OperationTypeTeardown
	OperationTypeInspection
	OperationTypeTransfer
)


//==============================================================
// WorkOrderStatus Declaration
//==============================================================
type WorkOrderStatus int
const (
    WorkOrderStatusPlanned WorkOrderStatus = iota
	WorkOrderStatusReleased
	WorkOrderStatusInProcess
	WorkOrderStatusHold
	WorkOrderStatusCompleted
	WorkOrderStatusClosed
	WorkOrderStatusCancelled
)


//==============================================================
// ScheduleStatus Declaration
//==============================================================
type ScheduleStatus int
const (
    ScheduleStatusDraft ScheduleStatus = iota
	ScheduleStatusApproved
	ScheduleStatusFrozen
	ScheduleStatusCompleted
)


//==============================================================
// SupplierTier Declaration
//==============================================================
type SupplierTier int
const (
    SupplierTierTier1 SupplierTier = iota
	SupplierTierTier2
	SupplierTierTier3
)


//==============================================================
// PurchaseOrderStatus Declaration
//==============================================================
type PurchaseOrderStatus int
const (
    PurchaseOrderStatusDraft PurchaseOrderStatus = iota
	PurchaseOrderStatusSubmitted
	PurchaseOrderStatusAcknowledged
	PurchaseOrderStatusPartiallyReceived
	PurchaseOrderStatusReceived
	PurchaseOrderStatusClosed
	PurchaseOrderStatusCancelled
)


//==============================================================
// ReceiptStatus Declaration
//==============================================================
type ReceiptStatus int
const (
    ReceiptStatusOpen ReceiptStatus = iota
	ReceiptStatusPartiallyProcessed
	ReceiptStatusCompleted
	ReceiptStatusRejected
)


//==============================================================
// WarehouseType Declaration
//==============================================================
type WarehouseType int
const (
    WarehouseTypeRawMaterial WarehouseType = iota
	WarehouseTypeWIP
	WarehouseTypeFinishedGoods
	WarehouseTypeDistribution
)


//==============================================================
// LocationType Declaration
//==============================================================
type LocationType int
const (
    LocationTypeBin LocationType = iota
	LocationTypeDock
	LocationTypeStaging
	LocationTypeQAHold
	LocationTypeScrap
)


//==============================================================
// InventoryTransactionType Declaration
//==============================================================
type InventoryTransactionType int
const (
    InventoryTransactionTypeReceipt InventoryTransactionType = iota
	InventoryTransactionTypeIssue
	InventoryTransactionTypeReturn
	InventoryTransactionTypeAdjustment
	InventoryTransactionTypeTransfer
	InventoryTransactionTypeConsumption
	InventoryTransactionTypeProductionReceipt
	InventoryTransactionTypeScrap
)


//==============================================================
// CustomerType Declaration
//==============================================================
type CustomerType int
const (
    CustomerTypeDistributor CustomerType = iota
	CustomerTypeOEM
	CustomerTypeRetailer
	CustomerTypeDirect
)


//==============================================================
// SalesOrderStatus Declaration
//==============================================================
type SalesOrderStatus int
const (
    SalesOrderStatusDraft SalesOrderStatus = iota
	SalesOrderStatusConfirmed
	SalesOrderStatusAllocated
	SalesOrderStatusInProduction
	SalesOrderStatusShipped
	SalesOrderStatusInvoiced
	SalesOrderStatusClosed
	SalesOrderStatusCancelled
)


//==============================================================
// SamplingPlanType Declaration
//==============================================================
type SamplingPlanType int
const (
    SamplingPlanTypeFixed SamplingPlanType = iota
	SamplingPlanTypePercentage
	SamplingPlanTypeC0
)


//==============================================================
// QualityPlanStatus Declaration
//==============================================================
type QualityPlanStatus int
const (
    QualityPlanStatusDraft QualityPlanStatus = iota
	QualityPlanStatusReleased
	QualityPlanStatusRetired
)


//==============================================================
// MeasurementType Declaration
//==============================================================
type MeasurementType int
const (
    MeasurementTypeAttribute MeasurementType = iota
	MeasurementTypeVariable
)


//==============================================================
// InspectionType Declaration
//==============================================================
type InspectionType int
const (
    InspectionTypeIncoming InspectionType = iota
	InspectionTypeInProcess
	InspectionTypeFinal
	InspectionTypeAudit
)


//==============================================================
// InspectionStatus Declaration
//==============================================================
type InspectionStatus int
const (
    InspectionStatusOpen InspectionStatus = iota
	InspectionStatusInProgress
	InspectionStatusCompleted
	InspectionStatusAccepted
	InspectionStatusRejected
)


//==============================================================
// InspectionResultStatus Declaration
//==============================================================
type InspectionResultStatus int
const (
    InspectionResultStatusPass InspectionResultStatus = iota
	InspectionResultStatusFail
	InspectionResultStatusRework
	InspectionResultStatusScrap
)


//==============================================================
// NonconformanceType Declaration
//==============================================================
type NonconformanceType int
const (
    NonconformanceTypeDimension NonconformanceType = iota
	NonconformanceTypeFunctional
	NonconformanceTypeCosmetic
	NonconformanceTypeDocumentation
	NonconformanceTypeSupplier
	NonconformanceTypeProcess
)


//==============================================================
// QualitySeverity Declaration
//==============================================================
type QualitySeverity int
const (
    QualitySeverityMinor QualitySeverity = iota
	QualitySeverityMajor
	QualitySeverityCritical
)


//==============================================================
// NonconformanceStatus Declaration
//==============================================================
type NonconformanceStatus int
const (
    NonconformanceStatusOpen NonconformanceStatus = iota
	NonconformanceStatusContained
	NonconformanceStatusUnderInvestigation
	NonconformanceStatusDispositioned
	NonconformanceStatusClosed
)


//==============================================================
// CAPAStatus Declaration
//==============================================================
type CAPAStatus int
const (
    CAPAStatusProposed CAPAStatus = iota
	CAPAStatusApproved
	CAPAStatusImplemented
	CAPAStatusVerified
	CAPAStatusClosed
)


//==============================================================
// AssetStatus Declaration
//==============================================================
type AssetStatus int
const (
    AssetStatusCommissioned AssetStatus = iota
	AssetStatusAvailable
	AssetStatusInMaintenance
	AssetStatusDown
	AssetStatusRetired
)


//==============================================================
// MaintenanceStrategy Declaration
//==============================================================
type MaintenanceStrategy int
const (
    MaintenanceStrategyTimeBased MaintenanceStrategy = iota
	MaintenanceStrategyUsageBased
	MaintenanceStrategyConditionBased
	MaintenanceStrategyPredictive
	MaintenanceStrategyCorrective
)


//==============================================================
// MaintenanceOrderStatus Declaration
//==============================================================
type MaintenanceOrderStatus int
const (
    MaintenanceOrderStatusCreated MaintenanceOrderStatus = iota
	MaintenanceOrderStatusApproved
	MaintenanceOrderStatusScheduled
	MaintenanceOrderStatusInProgress
	MaintenanceOrderStatusCompleted
	MaintenanceOrderStatusCancelled
)


//==============================================================
// EmployeeRole Declaration
//==============================================================
type EmployeeRole int
const (
    EmployeeRoleOperator EmployeeRole = iota
	EmployeeRoleTechnician
	EmployeeRoleSupervisor
	EmployeeRolePlanner
	EmployeeRoleQualityEngineer
	EmployeeRoleBuyer
)


//==============================================================
// SkillLevel Declaration
//==============================================================
type SkillLevel int
const (
    SkillLevelNovice SkillLevel = iota
	SkillLevelCompetent
	SkillLevelProficient
	SkillLevelExpert
)


//==============================================================
// ShiftType Declaration
//==============================================================
type ShiftType int
const (
    ShiftTypeDay ShiftType = iota
	ShiftTypeSwing
	ShiftTypeNight
	ShiftTypeWeekend
)


//==============================================================
// ForecastMethod Declaration
//==============================================================
type ForecastMethod int
const (
    ForecastMethodMovingAverage ForecastMethod = iota
	ForecastMethodExponentialSmoothing
	ForecastMethodCroston
	ForecastMethodARIMA
	ForecastMethodManual
)


//==============================================================
// MRPRunStatus Declaration
//==============================================================
type MRPRunStatus int
const (
    MRPRunStatusStarted MRPRunStatus = iota
	MRPRunStatusCompleted
	MRPRunStatusFailed
	MRPRunStatusCancelled
)


//==============================================================
// PlannedOrderType Declaration
//==============================================================
type PlannedOrderType int
const (
    PlannedOrderTypeWorkOrder PlannedOrderType = iota
	PlannedOrderTypePurchaseRequisition
	PlannedOrderTypeTransferOrder
)


//==============================================================
// PlannedOrderStatus Declaration
//==============================================================
type PlannedOrderStatus int
const (
    PlannedOrderStatusPlanned PlannedOrderStatus = iota
	PlannedOrderStatusFirmed
	PlannedOrderStatusReleased
	PlannedOrderStatusCancelled
)


//==============================================================
// PaymentTerms Declaration
//==============================================================
type PaymentTerms int
const (
    PaymentTermsNet30 PaymentTerms = iota
	PaymentTermsNet45
	PaymentTermsNet60
	PaymentTermsPrepaid
	PaymentTermsCOD
)

