
// Define collection and schema for TypeCertificate
export interface TypeCertificate {
    certificateNumber:
	type : string
    authority:
	type : string
    Program:
	type : Schema.Types.ObjectId
#
    collection: 'typeCertificates'
}
