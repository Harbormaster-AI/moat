
// Define collection and schema for Product
export interface Product {
    sku:
	type : string
    name:
	type : string
    asActive:
	type : boolean
    standardPrice:
	type : Money
    description:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    PriceBookEntries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PriceBookEntry' }]
    OpportunityLineItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OpportunityLineItem' }]
    QuoteLineItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'QuoteLineItem' }]
    OrderItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OrderItem' }]
    ProductType:
 	type : String
    Uom:
 	type : String
#
    collection: 'products'
}
