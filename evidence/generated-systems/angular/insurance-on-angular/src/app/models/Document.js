
// Define collection and schema for Document
export interface Document {
    fileName:
	type : string
    uploadedDate:
	type : Date
    Policy:
	type : Schema.Types.ObjectId
    Claim:
	type : Schema.Types.ObjectId
    Application:
	type : Schema.Types.ObjectId
    Customer:
	type : Schema.Types.ObjectId
    DocumentType:
 	type : String
#
    collection: 'documents'
}
