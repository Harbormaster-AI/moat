
// Define collection and schema for Issue
export interface Issue {
    title:
	type : string
    openedDate:
	type : Date
    closedDate:
	type : Date
    Risk:
	type : Schema.Types.ObjectId
    Finding:
	type : Schema.Types.ObjectId
    CorrectiveActions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CorrectiveAction' }]
    Control:
	type : Schema.Types.ObjectId
    IssueType:
 	type : String
    Priority:
 	type : String
    Status:
 	type : String
#
    collection: 'issues'
}
