
// Define collection and schema for InventorySource
export interface InventorySource {
    name:
	type : string
    domain:
	type : string
    Publisher:
	type : Schema.Types.ObjectId
    AdSlots:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AdSlot' }]
    Deals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Deal' }]
    Channel:
 	type : String
    PrimaryFormat:
 	type : String
#
    collection: 'inventorySources'
}
