
// Define collection and schema for Risk
export interface Risk {
    name:
	type : string
    description:
	type : string
    inherentRiskScore:
	type : number
    residualRiskScore:
	type : number
    Organization:
	type : Schema.Types.ObjectId
    Controls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Assessments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RiskAssessment' }]
    Issues:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Issue' }]
    Findings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditFinding' }]
    Category:
 	type : String
    Impact:
 	type : String
    Likelihood:
 	type : String
    Status:
 	type : String
#
    collection: 'risks'
}
