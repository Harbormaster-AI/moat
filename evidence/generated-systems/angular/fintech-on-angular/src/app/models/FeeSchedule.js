
// Define collection and schema for FeeSchedule
export interface FeeSchedule {
    name:
	type : string
    amount:
	type : Money
    percentage:
	type : String
    minimum:
	type : Money
    maximum:
	type : Money
    PricingPlan:
	type : Schema.Types.ObjectId
    FeeType:
 	type : String
    CalculationMethod:
 	type : String
#
    collection: 'feeSchedules'
}
