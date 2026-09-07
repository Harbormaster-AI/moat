
// Define collection and schema for FinancialInstitution
export interface FinancialInstitution {
    name:
	type : string
    legalName:
	type : string
    countryOfIncorporation:
	type : string
    bic:
	type : BIC
    website:
	type : string
    Branches:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Branch' }]
    Customers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Customer' }]
    ProductOfferings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ProductOffering' }]
    PaymentProcessors:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentProcessor' }]
    CompliancePolicies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CompliancePolicy' }]
#
    collection: 'financialInstitutions'
}
