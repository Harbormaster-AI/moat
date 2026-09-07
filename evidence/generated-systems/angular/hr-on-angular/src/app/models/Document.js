
// Define collection and schema for Document
export interface Document {
    name:
	type : string
    fileUrl:
	type : string
    uploadedDate:
	type : Date
    Candidate:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    DocumentType:
 	type : String
#
    collection: 'documents'
}
