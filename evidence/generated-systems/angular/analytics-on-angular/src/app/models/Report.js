
// Define collection and schema for Report
export interface Report {
    title:
	type : string
    audience:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    Visualizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Visualization' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    SemanticModels:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SemanticModel' }]
    Queries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BIQuery' }]
    Tags:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Tag' }]
    Status:
 	type : String
#
    collection: 'reports'
}
