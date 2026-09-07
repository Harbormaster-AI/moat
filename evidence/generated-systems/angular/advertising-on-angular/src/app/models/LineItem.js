
// Define collection and schema for LineItem
export interface LineItem {
    name:
	type : string
    bidAmount:
	type : Money
    dailyBudget:
	type : Money
    frequencyCap:
	type : FrequencyCap
    Campaign:
	type : Schema.Types.ObjectId
    Placements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Placement' }]
    TargetingProfile:
	type : Schema.Types.ObjectId
    Deal:
	type : Schema.Types.ObjectId
    Creatives:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CreativeAsset' }]
    PerformanceMetrics:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PerformanceMetric' }]
    Status:
 	type : String
    PricingModel:
 	type : String
    BidStrategy:
 	type : String
    Pacing:
 	type : String
#
    collection: 'lineItems'
}
