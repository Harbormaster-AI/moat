
// Define collection and schema for AuditFinding
export interface AuditFinding {
    title:
	type : string
    description:
	type : string
    dueDate:
	type : Date
    Engagement:
	type : Schema.Types.ObjectId
    Workpaper:
	type : Schema.Types.ObjectId
    CorrectiveActions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CorrectiveAction' }]
    RelatedRisks:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Risk' }]
    RelatedControls:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Control' }]
    Issues:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Issue' }]
    Severity:
 	type : String
    Status:
 	type : String
#
    collection: 'auditFindings'
}
