
// Define collection and schema for Wallet
export interface Wallet {
    currency:
	type : string
    balance:
	type : Money
    Customer:
	type : Schema.Types.ObjectId
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Transaction' }]
    Status:
 	type : String
#
    collection: 'wallets'
}
