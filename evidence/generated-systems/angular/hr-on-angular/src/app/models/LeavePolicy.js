
// Define collection and schema for LeavePolicy
export interface LeavePolicy {
    name:
	type : string
    accrualRate:
	type : String
    carryoverAllowed:
	type : boolean
    maxBalance:
	type : String
    Organization:
	type : Schema.Types.ObjectId
    LeaveRequests:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LeaveRequest' }]
    LeaveCategory:
 	type : String
    AccrualUnit:
 	type : String
#
    collection: 'leavePolicys'
}
