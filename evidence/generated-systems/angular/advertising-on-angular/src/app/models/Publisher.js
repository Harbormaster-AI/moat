
// Define collection and schema for Publisher
export interface Publisher {
    name:
	type : string
    website:
	type : string
    InventorySources:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InventorySource' }]
    Deals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Deal' }]
    CreativeApprovals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CreativeApproval' }]
    InsertionOrders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsertionOrder' }]
    RateCards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RateCard' }]
    PublisherType:
 	type : String
#
    collection: 'publishers'
}
