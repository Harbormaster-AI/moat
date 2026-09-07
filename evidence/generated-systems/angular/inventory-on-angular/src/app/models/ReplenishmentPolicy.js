
// Define collection and schema for ReplenishmentPolicy
export interface ReplenishmentPolicy {
    minLevel:
	type : String
    maxLevel:
	type : String
    reorderPoint:
	type : String
    reorderQuantity:
	type : String
    leadTimeDays:
	type : number
    reviewPeriodDays:
	type : number
    Sku:
	type : Schema.Types.ObjectId
    Warehouse:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    PolicyType:
 	type : String
#
    collection: 'replenishmentPolicys'
}
