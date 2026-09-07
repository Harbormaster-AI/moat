
// Define collection and schema for Feature
export interface Feature {
    name:
	type : string
    description:
	type : string
    FeatureSet:
	type : Schema.Types.ObjectId
    SourceDatasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    TrainingRuns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrainingRun' }]
    DataType:
 	type : String
#
    collection: 'features'
}
