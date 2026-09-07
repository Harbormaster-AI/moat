
// Define collection and schema for Team
export interface Team {
    name:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Users:
 	type : [{ type: Schema.Types.ObjectId, ref: 'User' }]
    Accounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    Opportunities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
    Cases:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Case_' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    TeamType:
 	type : String
#
    collection: 'teams'
}
