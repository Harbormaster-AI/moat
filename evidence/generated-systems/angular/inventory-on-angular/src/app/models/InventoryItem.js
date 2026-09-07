
// Define collection and schema for InventoryItem
export interface InventoryItem {
    quantityOnHand:
	type : String
    quantityAvailable:
	type : String
    quantityReserved:
	type : String
    unitCost:
	type : Money
    lastUpdated:
	type : Date
    Sku:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryTransaction' }]
    Reservations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Reservation' }]
    StockStatus:
 	type : String
#
    collection: 'inventoryItems'
}
