
// Define collection and schema for Contact
export interface Contact {
    firstName:
	type : string
    lastName:
	type : string
    title:
	type : string
    email:
	type : EmailAddress
    phone:
	type : PhoneNumber
    mobile:
	type : PhoneNumber
    mailingAddress:
	type : Address
    Organization:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    Opportunities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
    Cases:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Case_' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    Notes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Note' }]
    EmailMessages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmailMessage' }]
    PreferredContactMethod:
 	type : String
#
    collection: 'contacts'
}
