
// Define collection and schema for PlannedOrder
export interface PlannedOrder {
    plannedOrderNumber:
	type : string
    quantity:
	type : Quantity
    dueDate:
	type : Date
    MrpRun:
	type : Schema.Types.ObjectId
    Item:
	type : Schema.Types.ObjectId
    Plant:
	type : Schema.Types.ObjectId
    OrderType:
 	type : String
    Status:
 	type : String
#
    collection: 'plannedOrders'
}
