
// Define collection and schema for Creditor
export interface Creditor {
    name:
	type : string
    bic:
	type : BIC
    address:
	type : Address
    Mandates:
 	type : [{ type: Schema.Types.ObjectId, ref: 'DirectDebitMandate' }]
#
    collection: 'creditors'
}
