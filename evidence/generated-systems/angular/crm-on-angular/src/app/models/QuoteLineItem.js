
// Define collection and schema for QuoteLineItem
export interface QuoteLineItem {
    quantity:
	type : String
    unitPrice:
	type : Money
    discountAmount:
	type : Money
    taxAmount:
	type : Money
    totalAmount:
	type : Money
    Quote:
	type : Schema.Types.ObjectId
    Product:
	type : Schema.Types.ObjectId
    PriceBookEntry:
	type : Schema.Types.ObjectId
    OpportunityLineItem:
	type : Schema.Types.ObjectId
#
    collection: 'quoteLineItems'
}
