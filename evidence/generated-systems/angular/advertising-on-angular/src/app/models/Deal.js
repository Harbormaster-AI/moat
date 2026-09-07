
// Define collection and schema for Deal
export interface Deal {
    floorPrice:
	type : Money
    Publisher:
	type : Schema.Types.ObjectId
    InventorySources:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventorySource' }]
    Placements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Placement' }]
    DealType:
 	type : String
#
    collection: 'deals'
}
