
// Define collection and schema for Warehouse
export interface Warehouse {
    name:
	type : string
    code:
	type : string
    address:
	type : Address
    timeZone:
	type : string
    allowsOverAllocation:
	type : boolean
    StorageLocations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'StorageLocation' }]
    InventoryItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    InboundShipments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InboundShipment' }]
    OutboundAllocations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OutboundAllocation' }]
    OriginTransfers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TransferOrder' }]
    DestinationTransfers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TransferOrder' }]
    CycleCounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CycleCount' }]
#
    collection: 'warehouses'
}
