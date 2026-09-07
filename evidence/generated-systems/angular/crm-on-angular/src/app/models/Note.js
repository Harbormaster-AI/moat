
// Define collection and schema for Note
export interface Note {
    title:
	type : string
    content:
	type : string
    createdAt:
	type : Date
    updatedAt:
	type : Date
    Organization:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Contact:
	type : Schema.Types.ObjectId
    Opportunity:
	type : Schema.Types.ObjectId
    Case:
	type : Schema.Types.ObjectId
    Lead:
	type : Schema.Types.ObjectId
#
    collection: 'notes'
}
