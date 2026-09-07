
// Define collection and schema for Person
export interface Person {
    firstName:
	type : string
    lastName:
	type : string
    email:
	type : EmailAddress
    department:
	type : string
    RoleAssignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RoleAssignment' }]
    OwnedPolicies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    CorrectiveActions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CorrectiveAction' }]
#
    collection: 'persons'
}
