
// Define collection and schema for ComplianceAlert
export interface ComplianceAlert {
    alertCode:
	type : string
    raisedAt:
	type : DateTime
    notes:
	type : string
    Screening:
	type : Schema.Types.ObjectId
    Transaction:
	type : Schema.Types.ObjectId
    Severity:
 	type : String
    Status:
 	type : String
#
    collection: 'complianceAlerts'
}
