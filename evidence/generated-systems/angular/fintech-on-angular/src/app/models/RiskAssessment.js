
// Define collection and schema for RiskAssessment
export interface RiskAssessment {
    score:
	type : RiskScore
    assessedAt:
	type : DateTime
    modelVersion:
	type : string
    notes:
	type : string
    Application:
	type : Schema.Types.ObjectId
    Decision:
 	type : String
#
    collection: 'riskAssessments'
}
