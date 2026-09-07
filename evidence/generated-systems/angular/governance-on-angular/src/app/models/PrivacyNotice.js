
// Define collection and schema for PrivacyNotice
export interface PrivacyNotice {
    title:
	type : string
    audience:
	type : string
    versionLabel:
	type : string
    publicationDate:
	type : Date
    publicationUrl:
	type : URL
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    Organization:
	type : Schema.Types.ObjectId
    Consents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Consent' }]
    Status:
 	type : String
#
    collection: 'privacyNotices'
}
