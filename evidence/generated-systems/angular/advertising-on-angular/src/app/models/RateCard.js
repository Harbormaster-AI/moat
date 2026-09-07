
// Define collection and schema for RateCard
export interface RateCard {
    name:
	type : string
    effectiveDate:
	type : Date
    currency:
	type : string
    Publisher:
	type : Schema.Types.ObjectId
    Rates:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Rate' }]
#
    collection: 'rateCards'
}
