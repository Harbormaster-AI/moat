
// Define collection and schema for Coverage
export interface Coverage {
    memberId:
	type : string
    groupNumber:
	type : string
    effectiveDate:
	type : Date
    endDate:
	type : Date
    Patient:
	type : Schema.Types.ObjectId
    Plan:
	type : Schema.Types.ObjectId
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    Authorizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Authorization' }]
    CoverageType:
 	type : String
#
    collection: 'coverages'
}
