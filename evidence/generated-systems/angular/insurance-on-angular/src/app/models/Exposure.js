
// Define collection and schema for Exposure
export interface Exposure {
    Claim:
	type : Schema.Types.ObjectId
    PolicyCoverage:
	type : Schema.Types.ObjectId
    InsuredObject:
	type : Schema.Types.ObjectId
    Reserves:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ClaimReserve' }]
    Payments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ClaimPayment' }]
    ExposureType:
 	type : String
    Status:
 	type : String
#
    collection: 'exposures'
}
