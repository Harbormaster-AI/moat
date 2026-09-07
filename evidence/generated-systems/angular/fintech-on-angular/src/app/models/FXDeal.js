
// Define collection and schema for FXDeal
export interface FXDeal {
    dealReference:
	type : string
    baseCurrency:
	type : string
    quoteCurrency:
	type : string
    rate:
	type : String
    amount:
	type : Money
    settlementDate:
	type : Date
    Quote:
	type : Schema.Types.ObjectId
    PaymentOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentOrder' }]
    Status:
 	type : String
#
    collection: 'fXDeals'
}
