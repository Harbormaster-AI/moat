
// Define collection and schema for Interview
export interface Interview {
    interviewDate:
	type : Date
    feedback:
	type : string
    Requisition:
	type : Schema.Types.ObjectId
    Candidate:
	type : Schema.Types.ObjectId
    Interviewers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Employee' }]
    Stage:
 	type : String
    Result:
 	type : String
#
    collection: 'interviews'
}
