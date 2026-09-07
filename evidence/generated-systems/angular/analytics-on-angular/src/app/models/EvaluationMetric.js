
// Define collection and schema for EvaluationMetric
export interface EvaluationMetric {
    name:
	type : string
    value:
	type : String
    ModelVersion:
	type : Schema.Types.ObjectId
    Metric:
	type : Schema.Types.ObjectId
    Dataset:
	type : Schema.Types.ObjectId
#
    collection: 'evaluationMetrics'
}
