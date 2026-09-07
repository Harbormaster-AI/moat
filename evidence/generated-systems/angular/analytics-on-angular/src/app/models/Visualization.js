
// Define collection and schema for Visualization
export interface Visualization {
    title:
	type : string
    options:
	type : ChartOptions
    Dashboard:
	type : Schema.Types.ObjectId
    Report:
	type : Schema.Types.ObjectId
    Metrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Metric' }]
    Dimensions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dimension' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    ChartType:
 	type : String
#
    collection: 'visualizations'
}
