
// Define collection and schema for ComplianceProgram
export interface ComplianceProgram {
    name:
	type : string
    framework:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Requirements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ComplianceRequirement' }]
    Controls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Attestations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Attestation' }]
    Regulations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Regulation' }]
    Status:
 	type : String
#
    collection: 'compliancePrograms'
}
