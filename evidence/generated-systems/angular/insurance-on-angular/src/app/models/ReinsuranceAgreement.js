
// Define collection and schema for ReinsuranceAgreement
export interface ReinsuranceAgreement {
    agreementNumber:
	type : string
    effectivePeriod:
	type : DateRange
    retention:
	type : Money
    limit:
	type : Money
    cessionPercentage:
	type : Percentage
    Insurer:
	type : Schema.Types.ObjectId
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    ReinsuranceType:
 	type : String
    TreatyType:
 	type : String
#
    collection: 'reinsuranceAgreements'
}
