
// Define collection and schema for Claim
export interface Claim {
    claimNumber:
	type : ClaimNumber
    noticeDate:
	type : Date
    lossDate:
	type : Date
    reportedBy:
	type : string
    Policy:
	type : Schema.Types.ObjectId
    Customer:
	type : Schema.Types.ObjectId
    Adjuster:
	type : Schema.Types.ObjectId
    Incident:
	type : Schema.Types.ObjectId
    Exposures:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Exposure' }]
    Reserves:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ClaimReserve' }]
    ClaimPayments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ClaimPayment' }]
    ServiceProviders:
 	type : [{ type: Schema.Types.ObjectId, ref: 'ServiceProvider' }]
    Subrogations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SubrogationRecovery' }]
    Status:
 	type : String
    LossCause:
 	type : String
#
    collection: 'claims'
}
