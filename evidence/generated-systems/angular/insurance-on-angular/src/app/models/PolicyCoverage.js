
// Define collection and schema for PolicyCoverage
export interface PolicyCoverage {
    limit:
	type : Money
    deductible:
	type : Money
    premium:
	type : Money
    Policy:
	type : Schema.Types.ObjectId
    InsuredObjects:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsuredObject' }]
    CoverageType:
 	type : String
#
    collection: 'policyCoverages'
}
