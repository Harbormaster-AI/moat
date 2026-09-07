
// Define collection and schema for StorageLocation
export interface StorageLocation {
    code:
	type : string
    temperatureControlled:
	type : boolean
    capacity:
	type : String
    capacityUnit:
	type : string
    Warehouse:
	type : Schema.Types.ObjectId
    ParentLocation:
	type : Schema.Types.ObjectId
    ChildLocations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'StorageLocation' }]
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    LocationType:
 	type : String
#
    collection: 'storageLocations'
}
