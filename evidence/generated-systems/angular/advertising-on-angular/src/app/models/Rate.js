
// Define collection and schema for Rate
export interface Rate {
    unitPrice:
	type : Money
    RateCard:
	type : Schema.Types.ObjectId
    AdSlot:
	type : Schema.Types.ObjectId
    AdFormat:
 	type : String
    PricingModel:
 	type : String
#
    collection: 'rates'
}
