
// Define collection and schema for User
export interface User {
    username:
	type : string
    fullName:
	type : string
    email:
	type : EmailAddress
    locale:
	type : _Locale
    Organization:
	type : Schema.Types.ObjectId
    Teams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Team' }]
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    OwnedAccounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    OwnedLeads:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Lead' }]
    OwnedOpportunities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
    OwnedCases:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Case_' }]
    Quotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Quote' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Order' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contract' }]
    EmailMessages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmailMessage' }]
    Role:
 	type : String
    Status:
 	type : String
#
    collection: 'users'
}
