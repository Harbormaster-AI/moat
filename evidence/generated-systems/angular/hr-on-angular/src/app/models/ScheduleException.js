
// Define collection and schema for ScheduleException
export interface ScheduleException {
    date:
	type : Date
    reason:
	type : string
    hours:
	type : String
    WorkSchedule:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
#
    collection: 'scheduleExceptions'
}
