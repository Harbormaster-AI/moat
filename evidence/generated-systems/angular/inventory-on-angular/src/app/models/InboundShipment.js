
// Define collection and schema for InboundShipment
export interface InboundShipment {
    shipmentNumber:
	type : string
    expectedArrivalDate:
	type : Date
    arrivalDate:
	type : Date
    carrierName:
	type : string
    Warehouse:
	type : Schema.Types.ObjectId
    Lines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InboundShipmentLine' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryTransaction' }]
    Status:
 	type : String
#
    collection: 'inboundShipments'
}
