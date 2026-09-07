
// Define collection and schema for Record_
export interface Record_ {
    title:
	type : string
    creationDate:
	type : Date
    Repository:
	type : Schema.Types.ObjectId
    RetentionSchedule:
	type : Schema.Types.ObjectId
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    DataCategories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataCategory' }]
    LegalHolds:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LegalHold' }]
    DataSubjectRequests:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataSubjectRequest' }]
    RecordType:
 	type : String
    Classification:
 	type : String
    Status:
 	type : String
#
    collection: 'record_s'
}
