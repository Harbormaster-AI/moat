
// Define collection and schema for Policy
export interface Policy {
    title:
	type : string
    versionLabel:
	type : string
    approvalDate:
	type : Date
    nextReviewDate:
	type : Date
    documentUrl:
	type : URL
    Organization:
	type : Schema.Types.ObjectId
    Owners:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Person' }]
    RelatedRequirements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ComplianceRequirement' }]
    Controls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Procedures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Procedure' }]
    Exceptions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Exception_' }]
    Attestations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Attestation' }]
    PolicyType:
 	type : String
    Status:
 	type : String
#
    collection: 'policys'
}
