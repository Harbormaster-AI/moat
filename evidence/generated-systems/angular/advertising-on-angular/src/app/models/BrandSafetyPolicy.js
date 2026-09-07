
// Define collection and schema for BrandSafetyPolicy
export interface BrandSafetyPolicy {
    TargetingProfiles:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TargetingProfile' }]
    Level:
 	type : String
    ContentRatingThreshold:
 	type : String
#
    collection: 'brandSafetyPolicys'
}
