
// Define collection and schema for Certification
export interface Certification {
    name:
	type : string
    issuer:
	type : string
    validFrom:
	type : Date
    validTo:
	type : Date
    credentialId:
	type : string
    Employee:
	type : Schema.Types.ObjectId
    Course:
	type : Schema.Types.ObjectId
#
    collection: 'certifications'
}
