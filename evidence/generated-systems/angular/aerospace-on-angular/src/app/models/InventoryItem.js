
// Define collection and schema for InventoryItem
export interface InventoryItem {
    quantityOnHand:
	type : number
    quantityReserved:
	type : number
    lotNumber:
	type : string
    Component:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
#
    collection: 'inventoryItems'
}
