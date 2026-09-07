
// Define collection and schema for User
export interface User {
    firstName:
	type : string
    lastName:
	type : string
    email:
	type : Email
    Agency:
	type : Schema.Types.ObjectId
    Teams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Team' }]
    AdAccounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AdAccount' }]
    Role:
 	type : String
#
    collection: 'users'
}
