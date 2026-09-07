
// Define collection and schema for InvestmentPortfolio
export interface InvestmentPortfolio {
    portfolioCode:
	type : string
    baseCurrency:
	type : string
    createdAt:
	type : DateTime
    Customer:
	type : Schema.Types.ObjectId
    Accounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InvestmentAccount' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TradeOrder' }]
    Holdings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Position' }]
    Status:
 	type : String
#
    collection: 'investmentPortfolios'
}
