
// Define collection and schema for JobApplication
export interface JobApplication {
    applicationNumber:
	type : string
    appliedDate:
	type : Date
    resumeUrl:
	type : string
    Candidate:
	type : Schema.Types.ObjectId
    Requisition:
	type : Schema.Types.ObjectId
    Screenings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Screening' }]
    Status:
 	type : String
#
    collection: 'jobApplications'
}
