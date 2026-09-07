
// Define collection and schema for ModelVersion
export interface ModelVersion {
    version:
	type : string
    Model_:
	type : Schema.Types.ObjectId
    TrainingRun:
	type : Schema.Types.ObjectId
    EvaluationMetrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EvaluationMetric' }]
    Deployments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InferenceEndpoint' }]
    FeatureSets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FeatureSet' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Lifecycle:
 	type : String
    TrainingStatus:
 	type : String
#
    collection: 'modelVersions'
}
