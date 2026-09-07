
// Define collection and schema for Campaign
export interface Campaign {
    name:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    budget:
	type : Money
    actualCost:
	type : Money
    expectedRevenue:
	type : Money
    Organization:
	type : Schema.Types.ObjectId
    ParentCampaign:
	type : Schema.Types.ObjectId
    ChildCampaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    Members:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CampaignMember' }]
    Opportunities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Opportunity' }]
    Accounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    Leads:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Lead' }]
    Contacts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Contact' }]
    Teams:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Team' }]
    Activities:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Activity' }]
    Status:
 	type : String
    Type:
 	type : String
#
    collection: 'campaigns'
}
