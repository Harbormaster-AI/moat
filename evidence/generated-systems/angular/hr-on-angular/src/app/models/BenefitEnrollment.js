
// Define collection and schema for BenefitEnrollment
export interface BenefitEnrollment {
    enrollmentId:
	type : string
    effectiveFrom:
	type : Date
    effectiveTo:
	type : Date
    BenefitPlan:
	type : Schema.Types.ObjectId
    Employee:
	type : Schema.Types.ObjectId
    Dependents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dependent' }]
    Status:
 	type : String
    CoverageLevel:
 	type : String
#
    collection: 'benefitEnrollments'
}
