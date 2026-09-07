
// Define collection and schema for DataProcessingActivity
export interface DataProcessingActivity {
    name:
	type : string
    purpose:
	type : string
    startDate:
	type : Date
    Organization:
	type : Schema.Types.ObjectId
    DataCategories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataCategory' }]
    Systems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'System_' }]
    Records:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Record_' }]
    PrivacyNotices:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PrivacyNotice' }]
    ThirdParties:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ThirdParty' }]
    Consents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Consent' }]
    DataBreaches:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataBreach' }]
    DataSubjectRequests:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSubjectRequest' }]
    LawfulBasis:
 	type : String
#
    collection: 'dataProcessingActivitys'
}
