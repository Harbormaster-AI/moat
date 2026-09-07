
// Define collection and schema for WorkOrder
export interface WorkOrder {
    workOrderNumber:
	type : string
    plannedStart:
	type : Date
    plannedEnd:
	type : Date
    quantity:
	type : Quantity
    priority:
	type : number
    Item:
	type : Schema.Types.ObjectId
    Plant:
	type : Schema.Types.ObjectId
    Routing:
	type : Schema.Types.ObjectId
    Bom:
	type : Schema.Types.ObjectId
    ProductionSchedule:
	type : Schema.Types.ObjectId
    SalesOrder:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'workOrders'
}
