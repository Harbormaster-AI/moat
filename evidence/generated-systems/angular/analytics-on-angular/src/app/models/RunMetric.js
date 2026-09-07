
// Define collection and schema for RunMetric
export interface RunMetric {
    name:
	type : string
    value:
	type : String
    TrainingRun:
	type : Schema.Types.ObjectId
    Metric:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
#
    collection: 'runMetrics'
}
