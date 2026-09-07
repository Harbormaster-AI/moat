
// Define collection and schema for Obligation
export interface Obligation {
    referenceNumber:
	type : string
    descriptionText:
	type : string
    Regulation:
	type : Schema.Types.ObjectId
    Controls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contract' }]
    ObligationType:
 	type : String
    ReviewFrequency:
 	type : String
#
    collection: 'obligations'
}
