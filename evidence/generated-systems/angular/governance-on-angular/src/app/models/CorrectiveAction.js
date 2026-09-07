
// Define collection and schema for CorrectiveAction
export interface CorrectiveAction {
    actionTitle:
	type : string
    owner:
	type : string
    targetDate:
	type : Date
    Finding:
	type : Schema.Types.ObjectId
    Issue:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'correctiveActions'
}
