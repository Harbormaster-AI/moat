
// Define collection and schema for Dimension
export interface Dimension {
    name:
	type : string
    typeTime:
	type : boolean
    SemanticModel:
	type : Schema.Types.ObjectId
    Datasets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSet' }]
    GlossaryTerms:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessGlossaryTerm' }]
    DimensionType:
 	type : String
#
    collection: 'dimensions'
}
