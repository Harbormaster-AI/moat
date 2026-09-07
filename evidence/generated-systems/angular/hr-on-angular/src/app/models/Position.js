
// Define collection and schema for Position
export interface Position {
    positionCode:
	type : string
    fte:
	type : String
    Department:
	type : Schema.Types.ObjectId
    JobProfile:
	type : Schema.Types.ObjectId
    CostCenter:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    ManagerPosition:
	type : Schema.Types.ObjectId
    DirectReports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Position' }]
    Assignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmploymentAssignment' }]
    Status:
 	type : String
    WorkLocationType:
 	type : String
#
    collection: 'positions'
}
