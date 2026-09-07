
// Define collection and schema for AuditEngagement
export interface AuditEngagement {
    title:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    AuditProgram:
	type : Schema.Types.ObjectId
    BusinessUnits:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BusinessUnit' }]
    ControlTests:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ControlTest_' }]
    Workpapers:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditWorkpaper' }]
    Findings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditFinding' }]
    Status:
 	type : String
#
    collection: 'auditEngagements'
}
