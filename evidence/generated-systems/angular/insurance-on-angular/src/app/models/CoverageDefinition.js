
// Define collection and schema for CoverageDefinition
export interface CoverageDefinition {
    name:
	type : string
    defaultLimit:
	type : Money
    defaultDeductible:
	type : Money
    asMandatory:
	type : boolean
    Product:
	type : Schema.Types.ObjectId
    CoverageType:
 	type : String
#
    collection: 'coverageDefinitions'
}
