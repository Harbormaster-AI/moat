
// Define collection and schema for GovernanceBody
export interface GovernanceBody {
    name:
	type : string
    charterUrl:
	type : URL
    chair:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    RoleAssignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RoleAssignment' }]
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    BodyType:
 	type : String
#
    collection: 'governanceBodys'
}
