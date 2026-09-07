
// Define collection and schema for Consent
export interface Consent {
    grantedAt:
	type : DateTime
    expiresAt:
	type : DateTime
    scope:
	type : string
    Customer:
	type : Schema.Types.ObjectId
    ApiClient:
	type : Schema.Types.ObjectId
    ConsentType:
 	type : String
    Status:
 	type : String
#
    collection: 'consents'
}
