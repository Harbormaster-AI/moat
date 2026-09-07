
// Define collection and schema for Timesheet
export interface Timesheet {
    periodStart:
	type : Date
    periodEnd:
	type : Date
    submissionDate:
	type : Date
    Employee:
	type : Schema.Types.ObjectId
    TimeEntries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TimeEntry' }]
    Approvals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Approval' }]
    Status:
 	type : String
#
    collection: 'timesheets'
}
