
// Define collection and schema for Payout
export interface Payout {
    payoutReference:
	type : string
    amount:
	type : Money
    currency:
	type : string
    scheduledDate:
	type : Date
    paidDate:
	type : Date
    Merchant:
	type : Schema.Types.ObjectId
    SettlementBatch:
	type : Schema.Types.ObjectId
    DestinationAccount:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'payouts'
}
