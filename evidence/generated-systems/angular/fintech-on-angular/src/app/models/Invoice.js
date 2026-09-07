
// Define collection and schema for Invoice
export interface Invoice {
    invoiceNumber:
	type : string
    issueDate:
	type : Date
    dueDate:
	type : Date
    total:
	type : Money
    currency:
	type : string
    Merchant:
	type : Schema.Types.ObjectId
    Payments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentOrder' }]
    Status:
 	type : String
#
    collection: 'invoices'
}
