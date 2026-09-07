
// Define collection and schema for CycleCount
export interface CycleCount {
    countNumber:
	type : string
    scheduledDate:
	type : Date
    performedDate:
	type : Date
    approvedBy:
	type : string
    Warehouse:
	type : Schema.Types.ObjectId
    Locations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'StorageLocation' }]
    Entries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CycleCountEntry' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryTransaction' }]
    Status:
 	type : String
#
    collection: 'cycleCounts'
}
