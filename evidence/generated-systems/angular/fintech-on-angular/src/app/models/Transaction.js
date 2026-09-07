
// Define collection and schema for Transaction
export interface Transaction {
    amount:
	type : Money
    fee:
	type : Money
    exchangeRate:
	type : String
    createdAt:
	type : DateTime
    completedAt:
	type : DateTime
    narrative:
	type : string
    Account:
	type : Schema.Types.ObjectId
    Wallet:
	type : Schema.Types.ObjectId
    PaymentOrder:
	type : Schema.Types.ObjectId
    Merchant:
	type : Schema.Types.ObjectId
    Card:
	type : Schema.Types.ObjectId
    RelatedTransactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Transaction' }]
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ComplianceAlert' }]
    TransactionType:
 	type : String
    Status:
 	type : String
#
    collection: 'transactions'
}
