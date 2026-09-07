
// Define collection and schema for EmailMessage
export interface EmailMessage {
    subject:
	type : string
    body:
	type : string
    sentAt:
	type : Date
    messageId:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Contact:
	type : Schema.Types.ObjectId
    Lead:
	type : Schema.Types.ObjectId
    Case:
	type : Schema.Types.ObjectId
    Opportunity:
	type : Schema.Types.ObjectId
    Campaign:
	type : Schema.Types.ObjectId
    Direction:
 	type : String
    Status:
 	type : String
#
    collection: 'emailMessages'
}
