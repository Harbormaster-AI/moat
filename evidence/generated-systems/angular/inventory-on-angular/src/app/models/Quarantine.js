
// Define collection and schema for Quarantine
export interface Quarantine {
    reason:
	type : string
    startedAt:
	type : Date
    releasedAt:
	type : Date
    Warehouse:
	type : Schema.Types.ObjectId
    Items:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventoryItem' }]
    Lot:
	type : Schema.Types.ObjectId
    SerialNumbers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SerialNumber' }]
    Disposition:
 	type : String
#
    collection: 'quarantines'
}
