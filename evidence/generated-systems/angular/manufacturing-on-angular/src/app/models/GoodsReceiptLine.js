
// Define collection and schema for GoodsReceiptLine
export interface GoodsReceiptLine {
    lineNumber:
	type : number
    receivedQuantity:
	type : Quantity
    acceptedQuantity:
	type : Quantity
    rejectedQuantity:
	type : Quantity
    lot:
	type : LotId
    GoodsReceipt:
	type : Schema.Types.ObjectId
    Item:
	type : Schema.Types.ObjectId
    InventoryTransaction:
	type : Schema.Types.ObjectId
#
    collection: 'goodsReceiptLines'
}
