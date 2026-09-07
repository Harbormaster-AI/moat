
// Define collection and schema for CreativeApproval
export interface CreativeApproval {
    reviewer:
	type : string
    reviewedAt:
	type : Date
    CreativeAsset:
	type : Schema.Types.ObjectId
    Publisher:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'creativeApprovals'
}
