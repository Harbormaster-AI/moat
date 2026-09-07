
// Define collection and schema for Order
export interface Order {
    orderNumber:
	type : string
    orderDate:
	type : Date
    totalAmount:
	type : Money
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
    Quote:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Items:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OrderItem' }]
    Contract:
	type : Schema.Types.ObjectId
    PriceBook:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'orders'
}
