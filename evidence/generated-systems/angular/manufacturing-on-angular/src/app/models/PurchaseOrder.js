
// Define collection and schema for PurchaseOrder
export interface PurchaseOrder {
    poNumber:
	type : string
    orderDate:
	type : Date
    totalAmount:
	type : Money
    Supplier:
	type : Schema.Types.ObjectId
    Plant:
	type : Schema.Types.ObjectId
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PurchaseOrderLine' }]
    GoodsReceipts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'GoodsReceipt' }]
    Status:
 	type : String
#
    collection: 'purchaseOrders'
}
