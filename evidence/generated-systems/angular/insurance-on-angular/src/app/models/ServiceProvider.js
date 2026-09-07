
// Define collection and schema for ServiceProvider
export interface ServiceProvider {
    name:
	type : string
    taxId:
	type : string
    Claims:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Claim' }]
    ProviderType:
 	type : String
    NetworkStatus:
 	type : String
#
    collection: 'serviceProviders'
}
