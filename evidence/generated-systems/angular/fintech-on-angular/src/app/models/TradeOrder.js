
// Define collection and schema for TradeOrder
export interface TradeOrder {
    orderId:
	type : string
    quantity:
	type : String
    limitPrice:
	type : Money
    placedAt:
	type : DateTime
    Portfolio:
	type : Schema.Types.ObjectId
    Security:
	type : Schema.Types.ObjectId
    Trades:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Trade' }]
    Side:
 	type : String
    Type:
 	type : String
    Status:
 	type : String
    TimeInForce:
 	type : String
#
    collection: 'tradeOrders'
}
