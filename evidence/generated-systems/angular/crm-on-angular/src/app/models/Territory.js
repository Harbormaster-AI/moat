
// Define collection and schema for Territory
export interface Territory {
    name:
	type : string
    region:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Accounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    Users:
 	type : [{ type: Schema.Types.ObjectId, ref: 'User' }]
    TerritoryType:
 	type : String
#
    collection: 'territorys'
}
