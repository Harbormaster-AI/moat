
// Define collection and schema for ProductionSchedule
export interface ProductionSchedule {
    scheduleNumber:
	type : string
    horizonStart:
	type : Date
    horizonEnd:
	type : Date
    Plant:
	type : Schema.Types.ObjectId
    WorkOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'WorkOrder' }]
    Status:
 	type : String
#
    collection: 'productionSchedules'
}
