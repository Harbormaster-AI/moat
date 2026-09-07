
// enum type AircraftType
export let AircraftType = {
	NarrowBody:"NarrowBody",
	WideBody:"WideBody",
	RegionalJet:"RegionalJet",
	Turboprop:"Turboprop",
	BusinessJet:"BusinessJet",
	Helicopter:"Helicopter",
	eVTOL:"eVTOL",
	CargoPlane:"CargoPlane",
}

// enum type EngineCategory
export let EngineCategory = {
	Turbofan:"Turbofan",
	Turboprop:"Turboprop",
	Turbojet:"Turbojet",
	Piston:"Piston",
	Electric:"Electric",
	Rocket:"Rocket",
}

// enum type LandingGearType
export let LandingGearType = {
	Tricycle:"Tricycle",
	Tandem:"Tandem",
	Taildragger:"Taildragger",
	Skid:"Skid",
	Floats:"Floats",
	Retractable:"Retractable",
}

// enum type OptionCategory
export let OptionCategory = {
	Cabin:"Cabin",
	Connectivity:"Connectivity",
	Safety:"Safety",
	Performance:"Performance",
	Paint:"Paint",
	FlightDeck:"FlightDeck",
}

// enum type PackageType
export let PackageType = {
	PerformancePack:"PerformancePack",
	CabinPack:"CabinPack",
	ConnectivityPack:"ConnectivityPack",
	CompliancePack:"CompliancePack",
}

// enum type SupplierType
export let SupplierType = {
	Airframe:"Airframe",
	Engine:"Engine",
	Avionics:"Avionics",
	Systems:"Systems",
	Materials:"Materials",
	MRO:"MRO",
	Testing:"Testing",
}

// enum type SupplierApprovalStatus
export let SupplierApprovalStatus = {
	Applied:"Applied",
	Approved:"Approved",
	OnHold:"OnHold",
	Suspended:"Suspended",
}

// enum type ComponentCategory
export let ComponentCategory = {
	Structure:"Structure",
	System:"System",
	Avionics:"Avionics",
	Interior:"Interior",
	LandingGear:"LandingGear",
	Powerplant:"Powerplant",
	Consumable:"Consumable",
}

// enum type SerializationMethod
export let SerializationMethod = {
	Serialized:"Serialized",
	LotTracked:"LotTracked",
	None:"None",
}

// enum type ProductionLineType
export let ProductionLineType = {
	FinalAssembly:"FinalAssembly",
	SubAssembly:"SubAssembly",
	Integration:"Integration",
	TestAndDelivery:"TestAndDelivery",
}

// enum type ProductionOrderStatus
export let ProductionOrderStatus = {
	Planned:"Planned",
	Released:"Released",
	InAssembly:"InAssembly",
	FlightTest:"FlightTest",
	Completed:"Completed",
}

// enum type ScheduleStatus
export let ScheduleStatus = {
	Draft:"Draft",
	Published:"Published",
	Revised:"Revised",
	Closed:"Closed",
}

// enum type OperatorType
export let OperatorType = {
	Airline:"Airline",
	Cargo:"Cargo",
	Government:"Government",
	Private:"Private",
	Lessor:"Lessor",
}

// enum type AircraftOrderStatus
export let AircraftOrderStatus = {
	Draft:"Draft",
	Committed:"Committed",
	InProduction:"InProduction",
	Delivered:"Delivered",
	Cancelled:"Cancelled",
}

// enum type WarrantyType
export let WarrantyType = {
	Basic:"Basic",
	Powerplant:"Powerplant",
	Avionics:"Avionics",
	Corrosion:"Corrosion",
}

// enum type AppointmentStatus
export let AppointmentStatus = {
	Scheduled:"Scheduled",
	InProgress:"InProgress",
	Completed:"Completed",
	Cancelled:"Cancelled",
	Deferred:"Deferred",
}

// enum type WorkOrderStatus
export let WorkOrderStatus = {
	Open:"Open",
	InProgress:"InProgress",
	AwaitingParts:"AwaitingParts",
	Closed:"Closed",
	Deferred:"Deferred",
}

// enum type ServiceBulletinCategory
export let ServiceBulletinCategory = {
	Recommended:"Recommended",
	Optional:"Optional",
	Alert:"Alert",
	Mandatory:"Mandatory",
}

// enum type ConnectivityStatus
export let ConnectivityStatus = {
	Offline:"Offline",
	Online:"Online",
	Degraded:"Degraded",
}

// enum type EventSeverity
export let EventSeverity = {
	Info:"Info",
	Warning:"Warning",
	Critical:"Critical",
}

// enum type SoftwareLoadType
export let SoftwareLoadType = {
	FlightDeckSoftware:"FlightDeckSoftware",
	MaintenanceTools:"MaintenanceTools",
	CabinIFE:"CabinIFE",
	ConnectivityModem:"ConnectivityModem",
}

// enum type SalesCampaignStatus
export let SalesCampaignStatus = {
	Prospecting:"Prospecting",
	Proposal:"Proposal",
	Negotiation:"Negotiation",
	Won:"Won",
	Lost:"Lost",
}

// enum type ProgramStatus
export let ProgramStatus = {
	Concept:"Concept",
	Development:"Development",
	Certification:"Certification",
	Production:"Production",
	InService:"InService",
	Sunset:"Sunset",
}
