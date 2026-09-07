
// Define collection and schema for RoleAssignment
export interface RoleAssignment {
    effectiveFrom:
	type : Date
    effectiveTo:
	type : Date
    Person:
	type : Schema.Types.ObjectId
    Role:
	type : Schema.Types.ObjectId
    GovernanceBody:
	type : Schema.Types.ObjectId
    Organization:
	type : Schema.Types.ObjectId
#
    collection: 'roleAssignments'
}
