
// Define collection and schema for ComplianceRequirement
export interface ComplianceRequirement {
    name:
	type : string
    source:
	type : string
    citation:
	type : string
    ComplianceProgram:
	type : Schema.Types.ObjectId
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Controls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Obligations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Obligation' }]
    Applicability:
 	type : String
    Status:
 	type : String
#
    collection: 'complianceRequirements'
}
