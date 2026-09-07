
// Define collection and schema for Security
export interface Security {
    symbol:
	type : string
    isin:
	type : string
    cusip:
	type : string
    currency:
	type : string
    Positions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Position' }]
    Trades:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Trade' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TradeOrder' }]
    SecurityType:
 	type : String
#
    collection: 'securitys'
}
