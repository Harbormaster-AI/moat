
// Define collection and schema for Quote
export interface Quote {
    quoteNumber:
	type : string
    validityStart:
	type : Date
    validityEnd:
	type : Date
    totalAmount:
	type : Money
    discountPercent:
	type : String
    taxAmount:
	type : Money
    shippingAmount:
	type : Money
    Organization:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Opportunity:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    LineItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'QuoteLineItem' }]
    PriceBook:
	type : Schema.Types.ObjectId
    Order:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'quotes'
}
