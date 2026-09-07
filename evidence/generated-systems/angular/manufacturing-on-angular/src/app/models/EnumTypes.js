
// enum type BusinessUnitCategory
export let BusinessUnitCategory = {
	ConsumerGoods:"ConsumerGoods",
	IndustrialEquipment:"IndustrialEquipment",
	Electronics:"Electronics",
	Pharmaceuticals:"Pharmaceuticals",
	FoodBeverage:"FoodBeverage",
}

// enum type ProductionLineType
export let ProductionLineType = {
	Discrete:"Discrete",
	Batch:"Batch",
	Continuous:"Continuous",
	FlexibleCell:"FlexibleCell",
}

// enum type WorkCenterType
export let WorkCenterType = {
	Machining:"Machining",
	Assembly:"Assembly",
	Painting:"Painting",
	Packaging:"Packaging",
	Test:"Test",
	Warehouse:"Warehouse",
}

// enum type ItemType
export let ItemType = {
	FinishedGood:"FinishedGood",
	Subassembly:"Subassembly",
	Component:"Component",
	RawMaterial:"RawMaterial",
	Consumable:"Consumable",
	Service:"Service",
}

// enum type ProcurementType
export let ProcurementType = {
	MakeToStock:"MakeToStock",
	MakeToOrder:"MakeToOrder",
	Purchase:"Purchase",
	Kanban:"Kanban",
	Outsourced:"Outsourced",
}

// enum type UnitOfMeasure
export let UnitOfMeasure = {
	Each:"Each",
	Kilogram:"Kilogram",
	Gram:"Gram",
	Pound:"Pound",
	Liter:"Liter",
	Meter:"Meter",
	Centimeter:"Centimeter",
	Millimeter:"Millimeter",
	Hour:"Hour",
	Minute:"Minute",
	Box:"Box",
	Pallet:"Pallet",
}

// enum type TimeUnit
export let TimeUnit = {
	Second:"Second",
	Minute:"Minute",
	Hour:"Hour",
	Day:"Day",
}

// enum type ProductLifecycleStatus
export let ProductLifecycleStatus = {
	Active:"Active",
	PendingApproval:"PendingApproval",
	Discontinued:"Discontinued",
	Obsolete:"Obsolete",
}

// enum type BOMStatus
export let BOMStatus = {
	Draft:"Draft",
	Released:"Released",
	Obsolete:"Obsolete",
}

// enum type RoutingType
export let RoutingType = {
	Standard:"Standard",
	Alternate:"Alternate",
	Rework:"Rework",
}

// enum type RoutingStatus
export let RoutingStatus = {
	Draft:"Draft",
	Released:"Released",
	Obsolete:"Obsolete",
}

// enum type OperationType
export let OperationType = {
	Setup:"Setup",
	Run:"Run",
	Teardown:"Teardown",
	Inspection:"Inspection",
	Transfer:"Transfer",
}

// enum type WorkOrderStatus
export let WorkOrderStatus = {
	Planned:"Planned",
	Released:"Released",
	InProcess:"InProcess",
	Hold:"Hold",
	Completed:"Completed",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type ScheduleStatus
export let ScheduleStatus = {
	Draft:"Draft",
	Approved:"Approved",
	Frozen:"Frozen",
	Completed:"Completed",
}

// enum type SupplierTier
export let SupplierTier = {
	Tier1:"Tier1",
	Tier2:"Tier2",
	Tier3:"Tier3",
}

// enum type PurchaseOrderStatus
export let PurchaseOrderStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	Acknowledged:"Acknowledged",
	PartiallyReceived:"PartiallyReceived",
	Received:"Received",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type ReceiptStatus
export let ReceiptStatus = {
	Open:"Open",
	PartiallyProcessed:"PartiallyProcessed",
	Completed:"Completed",
	Rejected:"Rejected",
}

// enum type WarehouseType
export let WarehouseType = {
	RawMaterial:"RawMaterial",
	WIP:"WIP",
	FinishedGoods:"FinishedGoods",
	Distribution:"Distribution",
}

// enum type LocationType
export let LocationType = {
	Bin:"Bin",
	Dock:"Dock",
	Staging:"Staging",
	QAHold:"QAHold",
	Scrap:"Scrap",
}

// enum type InventoryTransactionType
export let InventoryTransactionType = {
	Receipt:"Receipt",
	Issue:"Issue",
	Return:"Return",
	Adjustment:"Adjustment",
	Transfer:"Transfer",
	Consumption:"Consumption",
	ProductionReceipt:"ProductionReceipt",
	Scrap:"Scrap",
}

// enum type CustomerType
export let CustomerType = {
	Distributor:"Distributor",
	OEM:"OEM",
	Retailer:"Retailer",
	Direct:"Direct",
}

// enum type SalesOrderStatus
export let SalesOrderStatus = {
	Draft:"Draft",
	Confirmed:"Confirmed",
	Allocated:"Allocated",
	InProduction:"InProduction",
	Shipped:"Shipped",
	Invoiced:"Invoiced",
	Closed:"Closed",
	Cancelled:"Cancelled",
}

// enum type SamplingPlanType
export let SamplingPlanType = {
	Fixed:"Fixed",
	Percentage:"Percentage",
	C0:"C0",
}

// enum type QualityPlanStatus
export let QualityPlanStatus = {
	Draft:"Draft",
	Released:"Released",
	Retired:"Retired",
}

// enum type MeasurementType
export let MeasurementType = {
	Attribute:"Attribute",
	Variable:"Variable",
}

// enum type InspectionType
export let InspectionType = {
	Incoming:"Incoming",
	InProcess:"InProcess",
	Final:"Final",
	Audit:"Audit",
}

// enum type InspectionStatus
export let InspectionStatus = {
	Open:"Open",
	InProgress:"InProgress",
	Completed:"Completed",
	Accepted:"Accepted",
	Rejected:"Rejected",
}

// enum type InspectionResultStatus
export let InspectionResultStatus = {
	Pass:"Pass",
	Fail:"Fail",
	Rework:"Rework",
	Scrap:"Scrap",
}

// enum type NonconformanceType
export let NonconformanceType = {
	Dimension:"Dimension",
	Functional:"Functional",
	Cosmetic:"Cosmetic",
	Documentation:"Documentation",
	Supplier:"Supplier",
	Process:"Process",
}

// enum type QualitySeverity
export let QualitySeverity = {
	Minor:"Minor",
	Major:"Major",
	Critical:"Critical",
}

// enum type NonconformanceStatus
export let NonconformanceStatus = {
	Open:"Open",
	Contained:"Contained",
	UnderInvestigation:"UnderInvestigation",
	Dispositioned:"Dispositioned",
	Closed:"Closed",
}

// enum type CAPAStatus
export let CAPAStatus = {
	Proposed:"Proposed",
	Approved:"Approved",
	Implemented:"Implemented",
	Verified:"Verified",
	Closed:"Closed",
}

// enum type AssetStatus
export let AssetStatus = {
	Commissioned:"Commissioned",
	Available:"Available",
	InMaintenance:"InMaintenance",
	Down:"Down",
	Retired:"Retired",
}

// enum type MaintenanceStrategy
export let MaintenanceStrategy = {
	TimeBased:"TimeBased",
	UsageBased:"UsageBased",
	ConditionBased:"ConditionBased",
	Predictive:"Predictive",
	Corrective:"Corrective",
}

// enum type MaintenanceOrderStatus
export let MaintenanceOrderStatus = {
	Created:"Created",
	Approved:"Approved",
	Scheduled:"Scheduled",
	InProgress:"InProgress",
	Completed:"Completed",
	Cancelled:"Cancelled",
}

// enum type EmployeeRole
export let EmployeeRole = {
	Operator:"Operator",
	Technician:"Technician",
	Supervisor:"Supervisor",
	Planner:"Planner",
	QualityEngineer:"QualityEngineer",
	Buyer:"Buyer",
}

// enum type SkillLevel
export let SkillLevel = {
	Novice:"Novice",
	Competent:"Competent",
	Proficient:"Proficient",
	Expert:"Expert",
}

// enum type ShiftType
export let ShiftType = {
	Day:"Day",
	Swing:"Swing",
	Night:"Night",
	Weekend:"Weekend",
}

// enum type ForecastMethod
export let ForecastMethod = {
	MovingAverage:"MovingAverage",
	ExponentialSmoothing:"ExponentialSmoothing",
	Croston:"Croston",
	ARIMA:"ARIMA",
	Manual:"Manual",
}

// enum type MRPRunStatus
export let MRPRunStatus = {
	Started:"Started",
	Completed:"Completed",
	Failed:"Failed",
	Cancelled:"Cancelled",
}

// enum type PlannedOrderType
export let PlannedOrderType = {
	WorkOrder:"WorkOrder",
	PurchaseRequisition:"PurchaseRequisition",
	TransferOrder:"TransferOrder",
}

// enum type PlannedOrderStatus
export let PlannedOrderStatus = {
	Planned:"Planned",
	Firmed:"Firmed",
	Released:"Released",
	Cancelled:"Cancelled",
}

// enum type PaymentTerms
export let PaymentTerms = {
	Net30:"Net30",
	Net45:"Net45",
	Net60:"Net60",
	Prepaid:"Prepaid",
	COD:"COD",
}
