
// Define collection and schema for PerformanceCycle
export interface PerformanceCycle {
    name:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    Organization:
	type : Schema.Types.ObjectId
    Reviews:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PerformanceReview' }]
    Goals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Goal' }]
    Status:
 	type : String
#
    collection: 'performanceCycles'
}
