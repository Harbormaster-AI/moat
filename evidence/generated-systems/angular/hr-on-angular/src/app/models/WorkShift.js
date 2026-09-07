
// Define collection and schema for WorkShift
export interface WorkShift {
    startTime:
	type : LocalTime
    endTime:
	type : LocalTime
    breakMinutes:
	type : number
    WorkSchedule:
	type : Schema.Types.ObjectId
    DayOfWeek:
 	type : String
#
    collection: 'workShifts'
}
