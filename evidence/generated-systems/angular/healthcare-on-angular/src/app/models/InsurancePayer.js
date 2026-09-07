
// Define collection and schema for InsurancePayer
export interface InsurancePayer {
    name:
	type : string
    website:
	type : string
    Plans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsurancePlan' }]
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    PayerType:
 	type : String
#
    collection: 'insurancePayers'
}
