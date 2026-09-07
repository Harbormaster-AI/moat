
// Define collection and schema for FeatureSet
export interface FeatureSet {
    name:
	type : string
    refreshSchedule:
	type : CronSchedule
    Workspace:
	type : Schema.Types.ObjectId
    Features:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Feature' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    ModelVersions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ModelVersion' }]
    Tags:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Tag' }]
    StoreType:
 	type : String
#
    collection: 'featureSets'
}
