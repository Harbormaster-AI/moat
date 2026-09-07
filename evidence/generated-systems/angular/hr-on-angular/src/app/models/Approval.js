
// Define collection and schema for Approval
export interface Approval {
    approverComment:
	type : string
    actionDate:
	type : Date
    Approver:
	type : Schema.Types.ObjectId
    Timesheet:
	type : Schema.Types.ObjectId
    LeaveRequest:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'approvals'
}
