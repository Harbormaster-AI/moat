
// Define collection and schema for BIQuery
export interface BIQuery {
    name:
	type : string
    text:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    Dashboards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dashboard' }]
    Notebooks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Notebook' }]
    Dialect:
 	type : String
#
    collection: 'bIQuerys'
}
