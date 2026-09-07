
// Define collection and schema for ShiftAssignment
export interface ShiftAssignment {
    assignmentDate:
	type : Date
    Shift:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    WorkCenter:
	type : Schema.Types.ObjectId
#
    collection: 'shiftAssignments'
}
