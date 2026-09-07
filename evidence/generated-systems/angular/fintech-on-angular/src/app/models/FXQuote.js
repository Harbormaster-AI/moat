
// Define collection and schema for FXQuote
export interface FXQuote {
    baseCurrency:
	type : string
    quoteCurrency:
	type : string
    rate:
	type : String
    quotedAt:
	type : DateTime
    expiresAt:
	type : DateTime
    RequestedBy:
	type : Schema.Types.ObjectId
    PriceType:
 	type : String
#
    collection: 'fXQuotes'
}
