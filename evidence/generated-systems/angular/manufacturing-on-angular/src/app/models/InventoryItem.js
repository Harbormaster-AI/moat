
// Define collection and schema for InventoryItem
export interface InventoryItem {
    quantityOnHand:
	type : Quantity
    quantityReserved:
	type : Quantity
    lotNumber:
	type : LotId
    serialNumber:
	type : SerialId
    Item:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
#
    collection: 'inventoryItems'
}
