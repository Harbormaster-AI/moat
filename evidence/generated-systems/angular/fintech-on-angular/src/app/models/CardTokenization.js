
// Define collection and schema for CardTokenization
export interface CardTokenization {
    tokenReference:
	type : string
    createdAt:
	type : DateTime
    Card:
	type : Schema.Types.ObjectId
    WalletProvider:
 	type : String
    Status:
 	type : String
#
    collection: 'cardTokenizations'
}
