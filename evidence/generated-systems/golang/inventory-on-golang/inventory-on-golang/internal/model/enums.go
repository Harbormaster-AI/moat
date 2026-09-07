package model


//==============================================================
// ItemType Declaration
//==============================================================
type ItemType int
const (
    ItemTypeFinishedGood ItemType = iota
	ItemTypeComponent
	ItemTypeRawMaterial
	ItemTypePackaging
	ItemTypeSparePart
	ItemTypeConsumable
)


//==============================================================
// UnitOfMeasure Declaration
//==============================================================
type UnitOfMeasure int
const (
    UnitOfMeasureEach UnitOfMeasure = iota
	UnitOfMeasureCase
	UnitOfMeasurePallet
	UnitOfMeasureDozen
	UnitOfMeasureGram
	UnitOfMeasureKilogram
	UnitOfMeasurePound
	UnitOfMeasureOunce
	UnitOfMeasureMilliliter
	UnitOfMeasureLiter
	UnitOfMeasureCubicMeter
	UnitOfMeasureMeter
	UnitOfMeasureFoot
	UnitOfMeasureSquareMeter
)


//==============================================================
// StockStatus Declaration
//==============================================================
type StockStatus int
const (
    StockStatusAvailable StockStatus = iota
	StockStatusReserved
	StockStatusDamaged
	StockStatusHold
	StockStatusQuarantined
	StockStatusInTransit
	StockStatusPendingInspection
)


//==============================================================
// LocationType Declaration
//==============================================================
type LocationType int
const (
    LocationTypeBin LocationType = iota
	LocationTypeBulk
	LocationTypeStaging
	LocationTypeDock
	LocationTypePicking
	LocationTypePacking
	LocationTypeQuality
	LocationTypeReturn
	LocationTypeColdStorage
)


//==============================================================
// LotStatus Declaration
//==============================================================
type LotStatus int
const (
    LotStatusReleased LotStatus = iota
	LotStatusQuarantined
	LotStatusExpired
	LotStatusBlocked
	LotStatusPendingTest
)


//==============================================================
// SerialStatus Declaration
//==============================================================
type SerialStatus int
const (
    SerialStatusActive SerialStatus = iota
	SerialStatusAssigned
	SerialStatusInTransit
	SerialStatusConsumed
	SerialStatusReturned
	SerialStatusScrapped
)


//==============================================================
// ReservationStatus Declaration
//==============================================================
type ReservationStatus int
const (
    ReservationStatusDraft ReservationStatus = iota
	ReservationStatusConfirmed
	ReservationStatusReleased
	ReservationStatusFulfilled
	ReservationStatusCancelled
	ReservationStatusExpired
)


//==============================================================
// ReservationType Declaration
//==============================================================
type ReservationType int
const (
    ReservationTypeSalesOrder ReservationType = iota
	ReservationTypeWorkOrder
	ReservationTypeTransferOrder
	ReservationTypeServiceOrder
	ReservationTypeOther
)


//==============================================================
// DemandType Declaration
//==============================================================
type DemandType int
const (
    DemandTypeSalesOrder DemandType = iota
	DemandTypeWorkOrder
	DemandTypeTransferOrder
	DemandTypeForecast
	DemandTypeSampleRequest
)


//==============================================================
// TransactionType Declaration
//==============================================================
type TransactionType int
const (
    TransactionTypeReceipt TransactionType = iota
	TransactionTypeIssue
	TransactionTypeAdjustmentIncrease
	TransactionTypeAdjustmentDecrease
	TransactionTypeReclassification
	TransactionTypeTransferOut
	TransactionTypeTransferIn
	TransactionTypeCountIncrease
	TransactionTypeCountDecrease
	TransactionTypePutaway
	TransactionTypePick
)


//==============================================================
// TransactionStatus Declaration
//==============================================================
type TransactionStatus int
const (
    TransactionStatusPending TransactionStatus = iota
	TransactionStatusPosted
	TransactionStatusVoided
)


//==============================================================
// TransferOrderStatus Declaration
//==============================================================
type TransferOrderStatus int
const (
    TransferOrderStatusDraft TransferOrderStatus = iota
	TransferOrderStatusReleased
	TransferOrderStatusInTransit
	TransferOrderStatusReceived
	TransferOrderStatusClosed
	TransferOrderStatusCancelled
)


//==============================================================
// AdjustmentType Declaration
//==============================================================
type AdjustmentType int
const (
    AdjustmentTypeIncrease AdjustmentType = iota
	AdjustmentTypeDecrease
	AdjustmentTypeReclassification
)


//==============================================================
// AdjustmentStatus Declaration
//==============================================================
type AdjustmentStatus int
const (
    AdjustmentStatusDraft AdjustmentStatus = iota
	AdjustmentStatusApproved
	AdjustmentStatusPosted
	AdjustmentStatusCancelled
)


//==============================================================
// CountStatus Declaration
//==============================================================
type CountStatus int
const (
    CountStatusPlanned CountStatus = iota
	CountStatusInProgress
	CountStatusCompleted
	CountStatusPosted
	CountStatusCancelled
)


//==============================================================
// ReplenishmentPolicyType Declaration
//==============================================================
type ReplenishmentPolicyType int
const (
    ReplenishmentPolicyTypeMinMax ReplenishmentPolicyType = iota
	ReplenishmentPolicyTypeReorderPoint
	ReplenishmentPolicyTypeEOQ
	ReplenishmentPolicyTypeKanban
)


//==============================================================
// InventoryAlertType Declaration
//==============================================================
type InventoryAlertType int
const (
    InventoryAlertTypeBelowMin InventoryAlertType = iota
	InventoryAlertTypeAboveMax
	InventoryAlertTypeStockoutRisk
	InventoryAlertTypeExcessStock
	InventoryAlertTypeExpiryRisk
)


//==============================================================
// AlertStatus Declaration
//==============================================================
type AlertStatus int
const (
    AlertStatusNew AlertStatus = iota
	AlertStatusAcknowledged
	AlertStatusResolved
	AlertStatusDismissed
)


//==============================================================
// Disposition Declaration
//==============================================================
type Disposition int
const (
    DispositionRelease Disposition = iota
	DispositionScrap
	DispositionReturnToVendor
	DispositionRework
)


//==============================================================
// RotationMethod Declaration
//==============================================================
type RotationMethod int
const (
    RotationMethodFIFO RotationMethod = iota
	RotationMethodLIFO
	RotationMethodFEFO
)


//==============================================================
// InboundShipmentStatus Declaration
//==============================================================
type InboundShipmentStatus int
const (
    InboundShipmentStatusPlanned InboundShipmentStatus = iota
	InboundShipmentStatusArrived
	InboundShipmentStatusReceived
	InboundShipmentStatusClosed
	InboundShipmentStatusCancelled
)


//==============================================================
// AllocationStatus Declaration
//==============================================================
type AllocationStatus int
const (
    AllocationStatusProposed AllocationStatus = iota
	AllocationStatusConfirmed
	AllocationStatusPicked
	AllocationStatusShort
	AllocationStatusCancelled
)

