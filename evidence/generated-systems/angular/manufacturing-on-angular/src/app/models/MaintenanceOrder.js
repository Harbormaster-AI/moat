
// Define collection and schema for MaintenanceOrder
export interface MaintenanceOrder {
    orderNumber:
	type : string
    priority:
	type : number
    requestedDate:
	type : Date
    completionDate:
	type : Date
    Asset:
	type : Schema.Types.ObjectId
    Plan:
	type : Schema.Types.ObjectId
    WorkCenter:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'maintenanceOrders'
}
