
// Define collection and schema for PricingPlan
export interface PricingPlan {
    name:
	type : string
    planCode:
	type : string
    baseCurrency:
	type : string
    ProductOffering:
	type : Schema.Types.ObjectId
    FeeSchedules:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FeeSchedule' }]
    Limits:
 	type : [{ type: Schema.Types.ObjectId, ref: 'UsageLimit' }]
    Status:
 	type : String
#
    collection: 'pricingPlans'
}
