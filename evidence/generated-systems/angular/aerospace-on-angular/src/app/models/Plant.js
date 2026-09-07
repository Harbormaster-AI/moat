
// Define collection and schema for Plant
export interface Plant {
    name:
	type : string
    plantCode:
	type : string
    address:
	type : Address
    Manufacturer:
	type : Schema.Types.ObjectId
    ProductionLines:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProductionLine' }]
    Warehouses:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Warehouse' }]
#
    collection: 'plants'
}
