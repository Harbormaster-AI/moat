
// Define collection and schema for BusinessGlossaryTerm
export interface BusinessGlossaryTerm {
    term:
	type : string
    definition:
	type : string
    steward:
	type : string
    RelatedTerms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessGlossaryTerm' }]
    Metrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Metric' }]
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    Dimensions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dimension' }]
    Measures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Measure' }]
#
    collection: 'businessGlossaryTerms'
}
