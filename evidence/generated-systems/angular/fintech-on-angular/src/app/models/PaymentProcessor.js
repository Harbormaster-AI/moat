
// Define collection and schema for PaymentProcessor
export interface PaymentProcessor {
    name:
	type : string
    processorCode:
	type : string
    networkSupport:
	type : string
    Institutions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'FinancialInstitution' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentContract' }]
    Settlements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SettlementBatch' }]
#
    collection: 'paymentProcessors'
}
