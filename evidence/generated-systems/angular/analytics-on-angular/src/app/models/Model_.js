
// Define collection and schema for Model_
export interface Model_ {
    name:
	type : string
    taskDescription:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    Versions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ModelVersion' }]
    FeatureSets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FeatureSet' }]
    Experiments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Experiment' }]
    Tags:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Tag' }]
    ModelType:
 	type : String
#
    collection: 'model_s'
}
