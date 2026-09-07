
// Define collection and schema for BusinessUnit
export interface BusinessUnit {
    name:
	type : string
    leader:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    Audits:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AuditEngagement' }]
#
    collection: 'businessUnits'
}
