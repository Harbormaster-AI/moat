
// Define collection and schema for Agent
export interface Agent {
    firstName:
	type : string
    lastName:
	type : string
    licenseId:
	type : string
    Distributor:
	type : Schema.Types.ObjectId
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Customers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Customer' }]
    Status:
 	type : String
#
    collection: 'agents'
}
