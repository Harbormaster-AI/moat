
// Define collection and schema for SubrogationRecovery
export interface SubrogationRecovery {
    recoveryReference:
	type : string
    amount:
	type : Money
    recoveryDate:
	type : Date
    Claim:
	type : Schema.Types.ObjectId
    Exposure:
	type : Schema.Types.ObjectId
    Counterparty:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'subrogationRecoverys'
}
