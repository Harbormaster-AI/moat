
// Define collection and schema for Routing
export interface Routing {
    routingNumber:
	type : string
    revision:
	type : string
    effectivityStart:
	type : Date
    effectivityEnd:
	type : Date
    Item:
	type : Schema.Types.ObjectId
    Operations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Operation' }]
    RoutingType:
 	type : String
    Status:
 	type : String
#
    collection: 'routings'
}
