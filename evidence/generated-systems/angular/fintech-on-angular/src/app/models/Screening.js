
// Define collection and schema for Screening
export interface Screening {
    score:
	type : RiskScore
    screenedAt:
	type : DateTime
    KycProfile:
	type : Schema.Types.ObjectId
    Alerts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ComplianceAlert' }]
    ScreeningType:
 	type : String
    Status:
 	type : String
#
    collection: 'screenings'
}
