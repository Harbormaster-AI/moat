
// Define collection and schema for Lead
export interface Lead {
    firstName:
	type : string
    lastName:
	type : string
    company:
	type : string
    email:
	type : EmailAddress
    phone:
	type : PhoneNumber
    converted:
	type : boolean
    Organization:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    ConvertedAccount:
	type : Schema.Types.ObjectId
    ConvertedContact:
	type : Schema.Types.ObjectId
    ConvertedOpportunity:
	type : Schema.Types.ObjectId
    Notes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Note' }]
    EmailMessages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmailMessage' }]
    Status:
 	type : String
    Source:
 	type : String
    Rating:
 	type : String
#
    collection: 'leads'
}
