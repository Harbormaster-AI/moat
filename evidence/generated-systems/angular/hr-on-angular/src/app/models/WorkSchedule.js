
// Define collection and schema for WorkSchedule
export interface WorkSchedule {
    name:
	type : string
    standardHoursPerWeek:
	type : String
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmploymentContract' }]
    Shifts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'WorkShift' }]
    Exceptions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ScheduleException' }]
    ScheduleType:
 	type : String
#
    collection: 'workSchedules'
}
