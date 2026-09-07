
// Define collection and schema for AuditProgram
export interface AuditProgram {
    name:
	type : string
    scope:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Engagements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditEngagement' }]
    Cycle:
 	type : String
    Status:
 	type : String
#
    collection: 'auditPrograms'
}
