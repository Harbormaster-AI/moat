
// Define collection and schema for BenefitPlan
export interface BenefitPlan {
    name:
	type : string
    providerName:
	type : string
    employeeContributionRate:
	type : Percentage
    employerContributionRate:
	type : Percentage
    eligibilityRules:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Enrollments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BenefitEnrollment' }]
    BenefitType:
 	type : String
#
    collection: 'benefitPlans'
}
