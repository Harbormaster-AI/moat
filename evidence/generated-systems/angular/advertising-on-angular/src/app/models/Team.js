
// Define collection and schema for Team
export interface Team {
    name:
	type : string
    Agency:
	type : Schema.Types.ObjectId
    Users:
 	type : [{ type: Schema.Types.ObjectId, ref: 'User' }]
    AdAccounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AdAccount' }]
#
    collection: 'teams'
}
