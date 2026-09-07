
// Define collection and schema for InsurancePlan
export interface InsurancePlan {
    name:
	type : string
    planCode:
	type : string
    Payer:
	type : Schema.Types.ObjectId
    Coverages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Coverage' }]
    PlanType:
 	type : String
#
    collection: 'insurancePlans'
}
