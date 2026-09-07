
// Define collection and schema for Endorsement
export interface Endorsement {
    endorsementNumber:
	type : string
    effectiveDate:
	type : Date
    description:
	type : string
    Policy:
	type : Schema.Types.ObjectId
#
    collection: 'endorsements'
}
