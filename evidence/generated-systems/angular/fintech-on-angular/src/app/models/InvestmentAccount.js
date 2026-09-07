
// Define collection and schema for InvestmentAccount
export interface InvestmentAccount {
    accountNumber:
	type : AccountNumber
    baseCurrency:
	type : string
    balance:
	type : Money
    Portfolio:
	type : Schema.Types.ObjectId
    Trades:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Trade' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TradeOrder' }]
    AccountType:
 	type : String
    Status:
 	type : String
#
    collection: 'investmentAccounts'
}
