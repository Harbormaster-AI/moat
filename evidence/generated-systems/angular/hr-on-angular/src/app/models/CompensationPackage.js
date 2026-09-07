
// Define collection and schema for CompensationPackage
export interface CompensationPackage {
    effectiveFrom:
	type : Date
    effectiveTo:
	type : Date
    currency:
	type : string
    Contract:
	type : Schema.Types.ObjectId
    SalaryComponents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SalaryComponent' }]
    BonusPlans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BonusPlan' }]
    EquityGrants:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EquityGrant' }]
#
    collection: 'compensationPackages'
}
