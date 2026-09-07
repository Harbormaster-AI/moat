
// Define collection and schema for CampaignMember
export interface CampaignMember {
    responded:
	type : boolean
    Campaign:
	type : Schema.Types.ObjectId
    Lead:
	type : Schema.Types.ObjectId
    Contact:
	type : Schema.Types.ObjectId
    Status:
 	type : String
    MemberType:
 	type : String
#
    collection: 'campaignMembers'
}
