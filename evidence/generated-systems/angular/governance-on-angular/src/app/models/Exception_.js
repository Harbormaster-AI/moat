
// Define collection and schema for Exception_
export interface Exception_ {
    title:
	type : string
    justification:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    RetentionSchedule:
	type : Schema.Types.ObjectId
    Policy:
	type : Schema.Types.ObjectId
    Control:
	type : Schema.Types.ObjectId
    Risk:
	type : Schema.Types.ObjectId
    ExceptionType:
 	type : String
    Status:
 	type : String
#
    collection: 'exception_s'
}
