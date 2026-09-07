
// Define collection and schema for RunParameter
export interface RunParameter {
    name:
	type : string
    value:
	type : string
    TrainingRun:
	type : Schema.Types.ObjectId
#
    collection: 'runParameters'
}
