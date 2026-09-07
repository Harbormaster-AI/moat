
// Define collection and schema for Shift
export interface Shift {
    shiftName:
	type : string
    startTime:
	type : string
    endTime:
	type : string
    Plant:
	type : Schema.Types.ObjectId
    Assignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ShiftAssignment' }]
    ShiftType:
 	type : String
#
    collection: 'shifts'
}
