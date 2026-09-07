
// enum type ItemType
export let ItemType = {
	FinishedGood:"FinishedGood",
	Component:"Component",
	RawMaterial:"RawMaterial",
	Packaging:"Packaging",
	SparePart:"SparePart",
	Consumable:"Consumable",
}

// enum type UnitOfMeasure
export let UnitOfMeasure = {
	Each:"Each",
	Case:"Case",
	Pallet:"Pallet",
	Dozen:"Dozen",
	Gram:"Gram",
	Kilogram:"Kilogram",
	Pound:"Pound",
	Ounce:"Ounce",
	Milliliter:"Milliliter",
	Liter:"Liter",
	CubicMeter:"CubicMeter",
	Meter:"Meter",
	Foot:"Foot",
	SquareMeter:"SquareMeter",
}

// enum type StockStatus
export let StockStatus = {
	Available:"Available",
	Reserved:"Reserved",
	Damaged:"Damaged",
	Hold:"Hold",
	Quarantined:"Quarantined",
	InTransit:"InTransit",
	PendingInspection:"PendingInspection",
}

// enum type LocationType
export let LocationType = {
	Bin:"Bin",
	Bulk:"Bulk",
	Staging:"Staging",
	Dock:"Dock",
	Picking:"Picking",
	Packing:"Packing",
	Quality:"Quality",
	Return:"Return",
	ColdStorage:"ColdStorage",
}

// enum type LotStatus
export let LotStatus = {
	Released:"Released",
	Quarantined:"Quarantined",
	Expired:"Expired",
	Blocked:"Blocked",
	PendingTest:"PendingTest",
}

// enum type SerialStatus
export let SerialStatus = {
	Active:"Active",
	Assigned:"Assigned",
	InTransit:"InTransit",
	Consumed:"Consumed",
	Returned:"Returned",
	Scrapped:"Scrapped",
}

// enum type ReservationStatus
export let ReservationStatus = {
	Draft:"Draft",
	Confirmed:"Confirmed",
	Released:"Released",
	Fulfilled:"Fulfilled",
	Cancelled:"Cancelled",
	Expired:"Expired",
}

// enum type ReservationType
export let ReservationType = {
	SalesOrder:"SalesOrder",
	WorkOrder:"WorkOrder",
	TransferOrder:"TransferOrder",
	ServiceOrder:"ServiceOrder",
	Other:"Other",
}

// enum type DemandType
export let DemandType = {
	SalesOrder:"SalesOrder",
	WorkOrder:"WorkOrder",
	TransferOrder:"TransferOrder",
	Forecast:"Forecast",
	SampleRequest:"SampleRequest",
}

// enum type TransactionType
export let TransactionType = {
	Receipt:"Receipt",
	Issue:"Issue",
	AdjustmentIncrease:"AdjustmentIncrease",
	AdjustmentDecrease:"AdjustmentDecrease",
	Reclassification:"Reclassification",
	TransferOut:"TransferOut",
	TransferIn:"TransferIn",
	CountIncrease:"CountIncrease",
	CountDecrease:"CountDecrease",
	Putaway:"Putaway",
	Pick:"Pick",
}

// enum type TransactionStatus
export let TransactionStatus = {
	Pending:"Pending",
	Posted:"Posted",
	Voided:"Voided",
}

// enum type TransferOrderStatus
export let TransferOrderStatus = {
	Draft:"Draft",
	Released:"Released",
	InTransit:"InTransit",
	Received:"Received",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type AdjustmentType
export let AdjustmentType = {
	Increase:"Increase",
	Decrease:"Decrease",
	Reclassification:"Reclassification",
}

// enum type AdjustmentStatus
export let AdjustmentStatus = {
	Draft:"Draft",
	Approved:"Approved",
	Posted:"Posted",
	Cancelled:"Cancelled",
}

// enum type CountStatus
export let CountStatus = {
	Planned:"Planned",
	InProgress:"InProgress",
	Completed:"Completed",
	Posted:"Posted",
	Cancelled:"Cancelled",
}

// enum type ReplenishmentPolicyType
export let ReplenishmentPolicyType = {
	MinMax:"MinMax",
	ReorderPoint:"ReorderPoint",
	EOQ:"EOQ",
	Kanban:"Kanban",
}

// enum type InventoryAlertType
export let InventoryAlertType = {
	BelowMin:"BelowMin",
	AboveMax:"AboveMax",
	StockoutRisk:"StockoutRisk",
	ExcessStock:"ExcessStock",
	ExpiryRisk:"ExpiryRisk",
}

// enum type AlertStatus
export let AlertStatus = {
	New:"New",
	Acknowledged:"Acknowledged",
	Resolved:"Resolved",
	Dismissed:"Dismissed",
}

// enum type Disposition
export let Disposition = {
	Release:"Release",
	Scrap:"Scrap",
	ReturnToVendor:"ReturnToVendor",
	Rework:"Rework",
}

// enum type RotationMethod
export let RotationMethod = {
	FIFO:"FIFO",
	LIFO:"LIFO",
	FEFO:"FEFO",
}

// enum type InboundShipmentStatus
export let InboundShipmentStatus = {
	Planned:"Planned",
	Arrived:"Arrived",
	Received:"Received",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type AllocationStatus
export let AllocationStatus = {
	Proposed:"Proposed",
	Confirmed:"Confirmed",
	Picked:"Picked",
	Short:"Short",
	Cancelled:"Cancelled",
}
