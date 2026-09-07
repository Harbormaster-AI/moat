
// Define collection and schema for EmploymentAssignment
export interface EmploymentAssignment {
    startDate:
	type : Date
    endDate:
	type : Date
    primary:
	type : boolean
    Employee:
	type : Schema.Types.ObjectId
    Position:
	type : Schema.Types.ObjectId
    Supervisor:
	type : Schema.Types.ObjectId
    AssignmentType:
 	type : String
    Status:
 	type : String
#
    collection: 'employmentAssignments'
}
