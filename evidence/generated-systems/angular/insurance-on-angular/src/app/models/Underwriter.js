
// Define collection and schema for Underwriter
export interface Underwriter {
    firstName:
	type : string
    lastName:
	type : string
    employeeId:
	type : string
    authorityLimit:
	type : Money
    Decisions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'UnderwritingDecision' }]
    Insurer:
	type : Schema.Types.ObjectId
#
    collection: 'underwriters'
}
