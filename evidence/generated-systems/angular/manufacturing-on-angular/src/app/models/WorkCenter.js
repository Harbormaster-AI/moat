
// Define collection and schema for WorkCenter
export interface WorkCenter {
    name:
	type : string
    code:
	type : string
    capacityPerHour:
	type : number
    oeeTarget:
	type : Percentage
    ProductionLine:
	type : Schema.Types.ObjectId
    Assets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Asset' }]
    MaintenanceOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'MaintenanceOrder' }]
    WorkCenterType:
 	type : String
#
    collection: 'workCenters'
}
