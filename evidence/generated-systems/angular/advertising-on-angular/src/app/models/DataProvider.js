
// Define collection and schema for DataProvider
export interface DataProvider {
    name:
	type : string
    website:
	type : string
    AudienceSegments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AudienceSegment' }]
    ProviderType:
 	type : String
#
    collection: 'dataProviders'
}
