
// Define collection and schema for MRPRun
export interface MRPRun {
    runNumber:
	type : string
    runDateTime:
	type : Date
    planningHorizonDays:
	type : number
    Plant:
	type : Schema.Types.ObjectId
    PlannedOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PlannedOrder' }]
    Status:
 	type : String
#
    collection: 'mRPRuns'
}
