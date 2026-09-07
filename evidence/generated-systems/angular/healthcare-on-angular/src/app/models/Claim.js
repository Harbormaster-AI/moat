
// Define collection and schema for Claim
export interface Claim {
    claimNumber:
	type : string
    totalAmount:
	type : Money
    Patient:
	type : Schema.Types.ObjectId
    Coverage:
	type : Schema.Types.ObjectId
    Encounter:
	type : Schema.Types.ObjectId
    Invoices:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Invoice' }]
    Payer:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'claims'
}
