
// Define collection and schema for Warehouse
export interface Warehouse {
    name:
	type : string
    warehouseCode:
	type : string
    address:
	type : Address
    Plant:
	type : Schema.Types.ObjectId
    Locations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Location' }]
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    WarehouseType:
 	type : String
#
    collection: 'warehouses'
}
