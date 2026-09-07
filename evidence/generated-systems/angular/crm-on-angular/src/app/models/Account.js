
// Define collection and schema for Account
export interface Account {
    name:
	type : string
    accountNumber:
	type : string
    industry:
	type : string
    billingAddress:
	type : Address
    shippingAddress:
	type : Address
    website:
	type : URL
    phone:
	type : PhoneNumber
    asActive:
	type : boolean
    Organization:
	type : Schema.Types.ObjectId
    ParentAccount:
	type : Schema.Types.ObjectId
    ChildAccounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    Contacts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contact' }]
    Opportunities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
    Cases:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Case_' }]
    Owner:
	type : Schema.Types.ObjectId
    Territory:
	type : Schema.Types.ObjectId
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    Quotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Quote' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Order' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contract' }]
    Notes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Note' }]
    EmailMessages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmailMessage' }]
    AccountType:
 	type : String
    LifecycleStage:
 	type : String
#
    collection: 'accounts'
}
