
// Define collection and schema for CompliancePolicy
export interface CompliancePolicy {
    name:
	type : string
    policyCode:
	type : string
    description:
	type : string
    Institution:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'compliancePolicys'
}
