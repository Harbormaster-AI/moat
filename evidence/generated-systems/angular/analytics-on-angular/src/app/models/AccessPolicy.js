
// Define collection and schema for AccessPolicy
export interface AccessPolicy {
    name:
	type : string
    subjectName:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Dashboards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dashboard' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    FeatureSets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FeatureSet' }]
    AccessLevel:
 	type : String
    SubjectType:
 	type : String
#
    collection: 'accessPolicys'
}
