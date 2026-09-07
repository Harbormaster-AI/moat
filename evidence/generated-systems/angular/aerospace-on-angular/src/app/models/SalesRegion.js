
// Define collection and schema for SalesRegion
export interface SalesRegion {
    name:
	type : string
    regionCode:
	type : string
    Operators:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Operator' }]
    SalesCampaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SalesCampaign' }]
#
    collection: 'salesRegions'
}
