
// Define collection and schema for Authorization
export interface Authorization {
    authNumber:
	type : string
    requestedService:
	type : string
    Coverage:
	type : Schema.Types.ObjectId
    Order:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'authorizations'
}
