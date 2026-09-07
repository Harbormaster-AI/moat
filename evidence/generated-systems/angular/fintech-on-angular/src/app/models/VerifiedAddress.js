
// Define collection and schema for VerifiedAddress
export interface VerifiedAddress {
    address:
	type : Address
    verifiedAt:
	type : DateTime
    KycProfile:
	type : Schema.Types.ObjectId
    VerificationStatus:
 	type : String
#
    collection: 'verifiedAddresss'
}
