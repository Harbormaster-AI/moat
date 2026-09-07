
// Define collection and schema for Opportunity
export interface Opportunity {
    name:
	type : string
    amount:
	type : Money
    closeDate:
	type : Date
    probability:
	type : String
    description:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Owner:
	type : Schema.Types.ObjectId
    Contacts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contact' }]
    LineItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OpportunityLineItem' }]
    StageHistory:
 	type : [{ type: Schema.Types.ObjectId, ref: 'OpportunityStageHistory' }]
    Quotes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Quote' }]
    Orders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Order' }]
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    Teams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Team' }]
    Stage:
 	type : String
    Type:
 	type : String
    ForecastCategory:
 	type : String
#
    collection: 'opportunitys'
}
