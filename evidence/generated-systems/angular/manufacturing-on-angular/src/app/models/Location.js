
// Define collection and schema for Location
export interface Location {
    locationCode:
	type : string
    description:
	type : string
    Warehouse:
	type : Schema.Types.ObjectId
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    LocationType:
 	type : String
#
    collection: 'locations'
}
