
// Define collection and schema for AnalyticsWorkspace
export interface AnalyticsWorkspace {
    name:
	type : string
    businessDomain:
	type : string
    ownerTeam:
	type : string
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    DataSources:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSource' }]
    Pipelines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataPipeline' }]
    Dashboards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dashboard' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    Notebooks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Notebook' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    FeatureSets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FeatureSet' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AccessPolicy' }]
    LineageNodes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LineageNode' }]
    GovernanceTier:
 	type : String
#
    collection: 'analyticsWorkspaces'
}
