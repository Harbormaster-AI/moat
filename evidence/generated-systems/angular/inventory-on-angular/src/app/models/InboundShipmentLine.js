
// Define collection and schema for InboundShipmentLine
export interface InboundShipmentLine {
    lineNumber:
	type : number
    quantity:
	type : String
    InboundShipment:
	type : Schema.Types.ObjectId
    Sku:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    DestinationLocation:
	type : Schema.Types.ObjectId
    UnitOfMeasure:
 	type : String
    StockStatus:
 	type : String
#
    collection: 'inboundShipmentLines'
}
