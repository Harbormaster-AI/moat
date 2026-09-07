
// Define collection and schema for LeaveRequest
export interface LeaveRequest {
    requestNumber:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    reason:
	type : string
    hours:
	type : String
    Employee:
	type : Schema.Types.ObjectId
    LeavePolicy:
	type : Schema.Types.ObjectId
    Approvals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Approval' }]
    Status:
 	type : String
#
    collection: 'leaveRequests'
}
