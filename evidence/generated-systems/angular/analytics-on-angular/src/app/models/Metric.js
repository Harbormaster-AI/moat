
// Define collection and schema for Metric
export interface Metric {
    name:
	type : string
    expression:
	type : string
    unit:
	type : string
    SemanticModel:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    GlossaryTerms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessGlossaryTerm' }]
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Alert' }]
    Visualizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Visualization' }]
    MetricType:
 	type : String
#
    collection: 'metrics'
}
