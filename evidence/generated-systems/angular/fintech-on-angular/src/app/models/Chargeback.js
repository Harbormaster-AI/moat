
// Define collection and schema for Chargeback
export interface Chargeback {
    chargebackReference:
	type : string
    amount:
	type : Money
    postedAt:
	type : DateTime
    Dispute:
	type : Schema.Types.ObjectId
    Transaction:
	type : Schema.Types.ObjectId
    Stage:
 	type : String
    Status:
 	type : String
#
    collection: 'chargebacks'
}
