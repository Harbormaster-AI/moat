
// Define collection and schema for UnderwritingDecision
export interface UnderwritingDecision {
    notes:
	type : string
    decisionDate:
	type : Date
    Quote:
	type : Schema.Types.ObjectId
    Underwriter:
	type : Schema.Types.ObjectId
    Decision:
 	type : String
#
    collection: 'underwritingDecisions'
}
