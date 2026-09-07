
// Define collection and schema for CycleCountEntry
export interface CycleCountEntry {
    lineNumber:
	type : number
    systemQuantity:
	type : String
    countedQuantity:
	type : String
    varianceQuantity:
	type : String
    recountRequired:
	type : boolean
    CycleCount:
	type : Schema.Types.ObjectId
    Sku:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    StockStatus:
 	type : String
#
    collection: 'cycleCountEntrys'
}
