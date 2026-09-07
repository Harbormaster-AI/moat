
// Define collection and schema for Dispute
export interface Dispute {
    disputeReference:
	type : string
    openedAt:
	type : DateTime
    closedAt:
	type : DateTime
    Transaction:
	type : Schema.Types.ObjectId
    Card:
	type : Schema.Types.ObjectId
    Merchant:
	type : Schema.Types.ObjectId
    Chargebacks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Chargeback' }]
    Reason:
 	type : String
    Status:
 	type : String
#
    collection: 'disputes'
}
