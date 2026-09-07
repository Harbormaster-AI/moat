
// Define collection and schema for Trade
export interface Trade {
    executedAt:
	type : DateTime
    quantity:
	type : String
    price:
	type : Money
    fees:
	type : Money
    settlementDate:
	type : Date
    Order:
	type : Schema.Types.ObjectId
    Security:
	type : Schema.Types.ObjectId
    InvestmentAccount:
	type : Schema.Types.ObjectId
#
    collection: 'trades'
}
