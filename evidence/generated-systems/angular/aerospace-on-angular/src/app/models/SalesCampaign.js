
// Define collection and schema for SalesCampaign
export interface SalesCampaign {
    campaignCode:
	type : string
    Region:
	type : Schema.Types.ObjectId
    Operator:
	type : Schema.Types.ObjectId
    Quotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Quote' }]
    Status:
 	type : String
#
    collection: 'salesCampaigns'
}
