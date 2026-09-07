
// Define collection and schema for TransferOrderLine
export interface TransferOrderLine {
    lineNumber:
	type : number
    quantity:
	type : String
    TransferOrder:
	type : Schema.Types.ObjectId
    Sku:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    FromLocation:
	type : Schema.Types.ObjectId
    ToLocation:
	type : Schema.Types.ObjectId
    UnitOfMeasure:
 	type : String
    StockStatus:
 	type : String
#
    collection: 'transferOrderLines'
}
