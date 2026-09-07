
// Define collection and schema for AuditWorkpaper
export interface AuditWorkpaper {
    workpaperRef:
	type : string
    subject:
	type : string
    workpaperUrl:
	type : URL
    Engagement:
	type : Schema.Types.ObjectId
    Evidence:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Evidence' }]
    Findings:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditFinding' }]
#
    collection: 'auditWorkpapers'
}
