
// Define collection and schema for SalesOrder
export interface SalesOrder {
    orderNumber:
	type : string
    orderDate:
	type : Date
    totalAmount:
	type : Money
    Customer:
	type : Schema.Types.ObjectId
    Plant:
	type : Schema.Types.ObjectId
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SalesOrderLine' }]
    WorkOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'WorkOrder' }]
    Status:
 	type : String
#
    collection: 'salesOrders'
}
