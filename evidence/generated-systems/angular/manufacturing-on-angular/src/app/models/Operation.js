
// Define collection and schema for Operation
export interface Operation {
    operationNumber:
	type : string
    name:
	type : string
    setupTime:
	type : TimeDuration
    standardCycleTime:
	type : TimeDuration
    Routing:
	type : Schema.Types.ObjectId
    WorkCenter:
	type : Schema.Types.ObjectId
    InspectionPlan:
	type : Schema.Types.ObjectId
    OperationType:
 	type : String
#
    collection: 'operations'
}
