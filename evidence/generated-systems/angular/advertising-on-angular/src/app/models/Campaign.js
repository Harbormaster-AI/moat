
// Define collection and schema for Campaign
export interface Campaign {
    name:
	type : string
    totalBudget:
	type : Money
    flight:
	type : DateRange
    AdAccount:
	type : Schema.Types.ObjectId
    LineItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LineItem' }]
    Kpis:
 	type : [{ type: Schema.Types.ObjectId, ref: 'KPI' }]
    TrackingPixels:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrackingPixel' }]
    Audiences:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AudienceSegment' }]
    Reports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Report' }]
    InsertionOrder:
	type : Schema.Types.ObjectId
    Objective:
 	type : String
    Status:
 	type : String
#
    collection: 'campaigns'
}
