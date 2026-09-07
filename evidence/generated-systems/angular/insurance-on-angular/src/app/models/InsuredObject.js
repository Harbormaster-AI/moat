
// Define collection and schema for InsuredObject
export interface InsuredObject {
    description:
	type : string
    serialOrId:
	type : string
    primaryAddress:
	type : Address
    Policy:
	type : Schema.Types.ObjectId
    Coverages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PolicyCoverage' }]
    ObjectType:
 	type : String
#
    collection: 'insuredObjects'
}
