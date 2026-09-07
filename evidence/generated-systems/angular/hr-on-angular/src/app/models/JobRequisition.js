
// Define collection and schema for JobRequisition
export interface JobRequisition {
    requisitionNumber:
	type : string
    title:
	type : string
    openings:
	type : number
    targetStartDate:
	type : Date
    Department:
	type : Schema.Types.ObjectId
    HiringManager:
	type : Schema.Types.ObjectId
    Recruiter:
	type : Schema.Types.ObjectId
    JobProfile:
	type : Schema.Types.ObjectId
    Candidates:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Candidate' }]
    Interviews:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Interview' }]
    Offers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Offer' }]
    Status:
 	type : String
    Priority:
 	type : String
#
    collection: 'jobRequisitions'
}
