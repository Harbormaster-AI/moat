
// Define collection and schema for KYCProfile
export interface KYCProfile {
    profileId:
	type : string
    createdAt:
	type : DateTime
    Customer:
	type : Schema.Types.ObjectId
    Documents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'KYCDocument' }]
    Screenings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Screening' }]
    Addresses:
 	type : [{ type: Schema.Types.ObjectId, ref: 'VerifiedAddress' }]
    Status:
 	type : String
    VerificationLevel:
 	type : String
#
    collection: 'kYCProfiles'
}
