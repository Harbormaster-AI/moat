
// Define collection and schema for Subscriber
export interface Subscriber {
    name:
	type : string
    address:
	type : string
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Alert' }]
    Channel:
 	type : String
#
    collection: 'subscribers'
}
