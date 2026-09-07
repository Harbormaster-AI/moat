
// Define collection and schema for TrainingRun
export interface TrainingRun {
    runLabel:
	type : string
    startedAt:
	type : Date
    completedAt:
	type : Date
    Experiment:
	type : Schema.Types.ObjectId
    ModelVersion:
	type : Schema.Types.ObjectId
    InputDatasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Features:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Feature' }]
    RunMetrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RunMetric' }]
    RunParameters:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RunParameter' }]
    Status:
 	type : String
#
    collection: 'trainingRuns'
}
