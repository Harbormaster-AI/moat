
// Define collection and schema for Warehouse
export interface Warehouse {
    name:
	type : string
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
#
    collection: 'warehouses'
}
