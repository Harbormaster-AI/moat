
// Define collection and schema for BuildSchedule
export interface BuildSchedule {
    scheduleNumber:
	type : string
    ProductionOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProductionOrder' }]
    Status:
 	type : String
#
    collection: 'buildSchedules'
}
