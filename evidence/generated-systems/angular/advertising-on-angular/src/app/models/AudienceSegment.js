
// Define collection and schema for AudienceSegment
export interface AudienceSegment {
    name:
	type : string
    estimatedReach:
	type : number
    description:
	type : string
    Provider:
	type : Schema.Types.ObjectId
    Campaigns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Campaign' }]
    ProviderType:
 	type : String
#
    collection: 'audienceSegments'
}
