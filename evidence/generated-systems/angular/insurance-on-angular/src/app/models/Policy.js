
// Define collection and schema for Policy
export interface Policy {
    policyNumber:
	type : PolicyNumber
    effectivePeriod:
	type : DateRange
    totalPremium:
	type : Money
    Insurer:
	type : Schema.Types.ObjectId
    Customer:
	type : Schema.Types.ObjectId
    Product:
	type : Schema.Types.ObjectId
    Agent:
	type : Schema.Types.ObjectId
    Coverages:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PolicyCoverage' }]
    InsuredObjects:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InsuredObject' }]
    Endorsements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Endorsement' }]
    BillingAccount:
	type : Schema.Types.ObjectId
    Beneficiaries:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Beneficiary' }]
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    ReinsuranceAgreements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ReinsuranceAgreement' }]
    Status:
 	type : String
    PaymentPlan:
 	type : String
#
    collection: 'policys'
}
