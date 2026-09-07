
// Define collection and schema for Measure
export interface Measure {
    name:
	type : string
    format:
	type : string
    SemanticModel:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    GlossaryTerms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessGlossaryTerm' }]
    Aggregation:
 	type : String
#
    collection: 'measures'
}
