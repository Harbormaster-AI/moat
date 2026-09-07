
// Define collection and schema for APIClient
export interface APIClient {
    name:
	type : string
    clientId:
	type : string
    redirectUri:
	type : string
    Consents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Consent' }]
    ClientType:
 	type : String
#
    collection: 'aPIClients'
}
