
// Define collection and schema for LineageNode
export interface LineageNode {
    name:
	type : string
    qualifiedName:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    Inputs:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LineageNode' }]
    Outputs:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LineageNode' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Models:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Model_' }]
    Pipelines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataPipeline' }]
    Dashboards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dashboard' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    NodeType:
 	type : String
#
    collection: 'lineageNodes'
}
