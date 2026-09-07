
// Define collection and schema for ThirdPartyAssessment
export interface ThirdPartyAssessment {
    assessmentDate:
	type : Date
    assessor:
	type : string
    ThirdParty:
	type : Schema.Types.ObjectId
    Issues:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Issue' }]
    AssessmentType:
 	type : String
    Result:
 	type : String
#
    collection: 'thirdPartyAssessments'
}
