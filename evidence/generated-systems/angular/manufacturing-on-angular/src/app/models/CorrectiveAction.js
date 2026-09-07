
// Define collection and schema for CorrectiveAction
export interface CorrectiveAction {
    capaNumber:
	type : string
    rootCause:
	type : string
    correctiveAction:
	type : string
    verificationDate:
	type : Date
    Nonconformance:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'correctiveActions'
}
