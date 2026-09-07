
// Define collection and schema for Employee
export interface Employee {
    firstName:
	type : string
    lastName:
	type : string
    WorkCenter:
	type : Schema.Types.ObjectId
    ShiftAssignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ShiftAssignment' }]
    CorrectiveActions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CorrectiveAction' }]
    Role:
 	type : String
    SkillLevel:
 	type : String
#
    collection: 'employees'
}
