
// Define collection and schema for InsertionOrder
export interface InsertionOrder {
    ioNumber:
	type : string
    agreedBudget:
	type : Money
    flight:
	type : DateRange
    Advertiser:
	type : Schema.Types.ObjectId
    Agency:
	type : Schema.Types.ObjectId
    Publisher:
	type : Schema.Types.ObjectId
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    Status:
 	type : String
#
    collection: 'insertionOrders'
}
