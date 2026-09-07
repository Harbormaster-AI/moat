
// Define collection and schema for StockAdjustment
export interface StockAdjustment {
    adjustmentNumber:
	type : string
    reason:
	type : string
    adjustmentDate:
	type : Date
    Warehouse:
	type : Schema.Types.ObjectId
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'StockAdjustmentLine' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryTransaction' }]
    AdjustmentType:
 	type : String
    Status:
 	type : String
#
    collection: 'stockAdjustments'
}
