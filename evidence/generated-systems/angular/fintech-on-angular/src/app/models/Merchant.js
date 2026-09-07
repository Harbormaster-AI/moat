
// Define collection and schema for Merchant
export interface Merchant {
    name:
	type : string
    mcc:
	type : string
    url:
	type : string
    country:
	type : string
    settlementCurrency:
	type : string
    Terminals:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Terminal' }]
    PaymentContracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentContract' }]
    Payouts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Payout' }]
    Settlements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'SettlementBatch' }]
    Disputes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dispute' }]
    Invoices:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Invoice' }]
#
    collection: 'merchants'
}
