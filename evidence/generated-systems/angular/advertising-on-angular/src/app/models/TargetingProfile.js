
// Define collection and schema for TargetingProfile
export interface TargetingProfile {
    name:
	type : string
    AudienceSegments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AudienceSegment' }]
    GeoRegions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'GeoRegion' }]
    ContentCategories:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ContentCategory' }]
    BrandSafetyPolicy:
	type : Schema.Types.ObjectId
    DeviceCriteria:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DeviceCriterion' }]
#
    collection: 'targetingProfiles'
}
