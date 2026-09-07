
// Define collection and schema for Dashboard
export interface Dashboard {
    title:
	type : string
    theme:
	type : string
    Workspace:
	type : Schema.Types.ObjectId
    Visualizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Visualization' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Alert' }]
    Queries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BIQuery' }]
    Tags:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Tag' }]
    Status:
 	type : String
#
    collection: 'dashboards'
}
