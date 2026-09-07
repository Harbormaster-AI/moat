
// Define collection and schema for Reservation
export interface Reservation {
    referenceNumber:
	type : string
    reservedQuantity:
	type : String
    promisedDate:
	type : Date
    Sku:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    InventoryItem:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    DemandSignal:
	type : Schema.Types.ObjectId
    ReservationStatus:
 	type : String
    ReservationType:
 	type : String
#
    collection: 'reservations'
}
