
// Define collection and schema for ProductOffering
export interface ProductOffering {
    name:
	type : string
    productCode:
	type : string
    Institution:
	type : Schema.Types.ObjectId
    PricingPlans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PricingPlan' }]
    Category:
 	type : String
#
    collection: 'productOfferings'
}
