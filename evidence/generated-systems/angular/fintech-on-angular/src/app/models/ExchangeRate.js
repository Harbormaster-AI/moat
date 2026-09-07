
// Define collection and schema for ExchangeRate
export interface ExchangeRate {
    baseCurrency:
	type : string
    quoteCurrency:
	type : string
    rate:
	type : String
    asOf:
	type : DateTime
    source:
	type : string
    UsedByQuotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FXQuote' }]
#
    collection: 'exchangeRates'
}
