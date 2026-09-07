
// Define collection and schema for InventoryItem
export interface InventoryItem {
    sku:
	type : string
    name:
	type : string
    quantityOnHand:
	type : number
    quantityReserved:
	type : number
    Facility:
	type : Schema.Types.ObjectId
    Supplier:
	type : Schema.Types.ObjectId
#
    collection: 'inventoryItems'
}
