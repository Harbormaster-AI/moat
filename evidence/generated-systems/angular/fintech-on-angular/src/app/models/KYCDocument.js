
// Define collection and schema for KYCDocument
export interface KYCDocument {
    reference:
	type : DocumentReference
    issuedCountry:
	type : string
    expirationDate:
	type : Date
    KycProfile:
	type : Schema.Types.ObjectId
    DocumentType:
 	type : String
    Status:
 	type : String
#
    collection: 'kYCDocuments'
}
