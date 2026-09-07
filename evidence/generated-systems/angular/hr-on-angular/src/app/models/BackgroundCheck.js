
// Define collection and schema for BackgroundCheck
export interface BackgroundCheck {
    checkNumber:
	type : string
    provider:
	type : string
    completedDate:
	type : Date
    Candidate:
	type : Schema.Types.ObjectId
    Requisition:
	type : Schema.Types.ObjectId
    Report:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'backgroundChecks'
}
