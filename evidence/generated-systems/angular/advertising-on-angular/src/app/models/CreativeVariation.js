
// Define collection and schema for CreativeVariation
export interface CreativeVariation {
    name:
	type : string
    language:
	type : string
    headline:
	type : string
    bodyText:
	type : string
    callToAction:
	type : string
    CreativeAsset:
	type : Schema.Types.ObjectId
#
    collection: 'creativeVariations'
}
