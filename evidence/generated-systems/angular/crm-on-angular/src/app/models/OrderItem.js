
// Define collection and schema for OrderItem
export interface OrderItem {
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
    Order:
	type : Schema.Types.ObjectId
    Product:
	type : Schema.Types.ObjectId
    PriceBookEntry:
	type : Schema.Types.ObjectId
#
    collection: 'orderItems'
}
