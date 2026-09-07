
// Define collection and schema for MaintenanceWorkOrder
export interface MaintenanceWorkOrder {
    workOrderNumber:
	type : string
    Aircraft:
	type : Schema.Types.ObjectId
    AirworthinessDirective:
	type : Schema.Types.ObjectId
    ServiceBulletin:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'maintenanceWorkOrders'
}
