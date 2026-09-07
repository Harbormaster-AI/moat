
// Define collection and schema for RetentionSchedule
export interface RetentionSchedule {
    name:
	type : string
    retentionPeriodMonths:
	type : number
    Repositories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RecordsRepository' }]
    Records:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Record_' }]
    Exceptions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Exception_' }]
    DispositionReviews:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DispositionReview' }]
    RetentionTrigger:
 	type : String
    DispositionAction:
 	type : String
    Status:
 	type : String
#
    collection: 'retentionSchedules'
}
