
// Define collection and schema for StockAdjustmentLine
export interface StockAdjustmentLine {
    lineNumber:
	type : number
    quantity:
	type : String
    Adjustment:
	type : Schema.Types.ObjectId
    Sku:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    UnitOfMeasure:
 	type : String
    StockStatus:
 	type : String
#
    collection: 'stockAdjustmentLines'
}
