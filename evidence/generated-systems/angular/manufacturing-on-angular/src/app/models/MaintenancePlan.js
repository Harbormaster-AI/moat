
// Define collection and schema for MaintenancePlan
export interface MaintenancePlan {
    planNumber:
	type : string
    interval:
	type : TimeDuration
    lastServiceDate:
	type : Date
    Asset:
	type : Schema.Types.ObjectId
    MaintenanceOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceOrder' }]
    Strategy:
 	type : String
#
    collection: 'maintenancePlans'
}
