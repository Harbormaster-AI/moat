
// Define collection and schema for UsageLimit
export interface UsageLimit {
    name:
	type : string
    amount:
	type : Money
    count:
	type : number
    PricingPlan:
	type : Schema.Types.ObjectId
    Scope:
 	type : String
    Period:
 	type : String
#
    collection: 'usageLimits'
}
