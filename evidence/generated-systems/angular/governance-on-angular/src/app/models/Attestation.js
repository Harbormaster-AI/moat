
// Define collection and schema for Attestation
export interface Attestation {
    statement:
	type : string
    attestor:
	type : string
    dateSigned:
	type : Date
    Control:
	type : Schema.Types.ObjectId
    Policy:
	type : Schema.Types.ObjectId
    ComplianceProgram:
	type : Schema.Types.ObjectId
    Result:
 	type : String
#
    collection: 'attestations'
}
