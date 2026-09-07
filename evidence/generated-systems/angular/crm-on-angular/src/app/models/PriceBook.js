
// Define collection and schema for PriceBook
export interface PriceBook {
    name:
	type : string
    asActive:
	type : boolean
    description:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Entries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PriceBookEntry' }]
    Quotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Quote' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Order' }]
#
    collection: 'priceBooks'
}
