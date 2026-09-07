
// Define collection and schema for DataSubjectRequest
export interface DataSubjectRequest {
    receivedDate:
	type : Date
    dueDate:
	type : Date
    requesterCountry:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    Records:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Record_' }]
    RequestType:
 	type : String
    Status:
 	type : String
#
    collection: 'dataSubjectRequests'
}
