
// Define collection and schema for Consent
export interface Consent {
    subjectIdentifier:
	type : string
    captureDate:
	type : Date
    expiryDate:
	type : Date
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    PrivacyNotice:
	type : Schema.Types.ObjectId
    ConsentType:
 	type : String
    Status:
 	type : String
#
    collection: 'consents'
}
