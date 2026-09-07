
// Define collection and schema for DataBreach
export interface DataBreach {
    incidentDate:
	type : Date
    description:
	type : string
    recordsAffected:
	type : number
    notificationRequired:
	type : boolean
    Organization:
	type : Schema.Types.ObjectId
    ProcessingActivities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataProcessingActivity' }]
    DataCategories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DataCategory' }]
    ThirdParties:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ThirdParty' }]
    Matter:
	type : Schema.Types.ObjectId
    Severity:
 	type : String
    Status:
 	type : String
#
    collection: 'dataBreachs'
}
