
// Define collection and schema for SettlementBatch
export interface SettlementBatch {
    batchId:
	type : string
    periodStart:
	type : DateTime
    periodEnd:
	type : DateTime
    totalVolume:
	type : Money
    totalCount:
	type : number
    Processor:
	type : Schema.Types.ObjectId
    Merchant:
	type : Schema.Types.ObjectId
    Payouts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Payout' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Transaction' }]
    Status:
 	type : String
#
    collection: 'settlementBatchs'
}
