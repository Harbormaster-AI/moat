
// Define collection and schema for OutboundAllocation
export interface OutboundAllocation {
    allocationNumber:
	type : string
    allocatedQuantity:
	type : String
    allocationDate:
	type : Date
    Warehouse:
	type : Schema.Types.ObjectId
    Sku:
	type : Schema.Types.ObjectId
    InventoryItem:
	type : Schema.Types.ObjectId
    Reservation:
	type : Schema.Types.ObjectId
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    SourceLocation:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'outboundAllocations'
}
