
// Define collection and schema for Condition
export interface Condition {
    code:
	type : string
    onsetDate:
	type : Date
    abatementDate:
	type : Date
    Patient:
	type : Schema.Types.ObjectId
    ClinicalStatus:
 	type : String
    VerificationStatus:
 	type : String
#
    collection: 'conditions'
}
