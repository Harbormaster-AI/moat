
// Define collection and schema for SemanticModel
export interface SemanticModel {
    name:
	type : string
    version:
	type : string
    grain:
	type : string
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Metrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Metric' }]
    Dimensions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dimension' }]
    Measures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Measure' }]
    GlossaryTerms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessGlossaryTerm' }]
#
    collection: 'semanticModels'
}
