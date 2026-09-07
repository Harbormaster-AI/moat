
// Define collection and schema for Termination
export interface Termination {
    terminationNumber:
	type : string
    terminationDate:
	type : Date
    notes:
	type : string
    eligibleForRehire:
	type : boolean
    Employee:
	type : Schema.Types.ObjectId
    Assignment:
	type : Schema.Types.ObjectId
    Reason:
 	type : String
    Type:
 	type : String
#
    collection: 'terminations'
}
