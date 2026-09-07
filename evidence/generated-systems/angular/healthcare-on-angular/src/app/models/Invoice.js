
// Define collection and schema for Invoice
export interface Invoice {
    invoiceNumber:
	type : string
    totalAmount:
	type : Money
    dueDate:
	type : Date
    Patient:
	type : Schema.Types.ObjectId
    Claim:
	type : Schema.Types.ObjectId
    Payments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Payment' }]
    Status:
 	type : String
#
    collection: 'invoices'
}
