
// Define collection and schema for Role
export interface Role {
    name:
	type : string
    responsibility:
	type : string
    Assignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RoleAssignment' }]
#
    collection: 'roles'
}
