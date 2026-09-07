
// Define collection and schema for GoodsReceipt
export interface GoodsReceipt {
    receiptNumber:
	type : string
    receiptDate:
	type : Date
    PurchaseOrder:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'GoodsReceiptLine' }]
    Status:
 	type : String
#
    collection: 'goodsReceipts'
}
