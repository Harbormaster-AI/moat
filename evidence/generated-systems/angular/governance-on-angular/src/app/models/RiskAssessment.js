
// Define collection and schema for RiskAssessment
export interface RiskAssessment {
    assessmentDate:
	type : Date
    assessor:
	type : string
    summary:
	type : string
    Risk:
	type : Schema.Types.ObjectId
    AssessmentType:
 	type : String
#
    collection: 'riskAssessments'
}
