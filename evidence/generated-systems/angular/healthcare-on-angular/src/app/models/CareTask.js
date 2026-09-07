
// Define collection and schema for CareTask
export interface CareTask {
    description:
	type : string
    dueDate:
	type : Date
    CarePlan:
	type : Schema.Types.ObjectId
    AssignedTo:
	type : Schema.Types.ObjectId
    Encounter:
	type : Schema.Types.ObjectId
    Status:
 	type : String
    Priority:
 	type : String
#
    collection: 'careTasks'
}
