
// Define collection and schema for Tag
export interface Tag {
    name:
	type : string
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    ModelVersions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ModelVersion' }]
    Dashboards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dashboard' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    FeatureSets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FeatureSet' }]
    Metrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Metric' }]
    Category:
 	type : String
#
    collection: 'tags'
}
