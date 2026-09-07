
// Define collection and schema for InsuranceProduct
export interface InsuranceProduct {
    name:
	type : string
    productCode:
	type : string
    Insurer:
	type : Schema.Types.ObjectId
    CoverageDefinitions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CoverageDefinition' }]
    LineOfBusiness:
 	type : String
#
    collection: 'insuranceProducts'
}
