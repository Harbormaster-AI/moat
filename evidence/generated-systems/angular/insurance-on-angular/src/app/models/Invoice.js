
// Define collection and schema for Invoice
export interface Invoice {
    invoiceNumber:
	type : string
    dueDate:
	type : Date
    totalDue:
	type : Money
    BillingAccount:
	type : Schema.Types.ObjectId
    Policy:
	type : Schema.Types.ObjectId
    Payments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Payment' }]
    Status:
 	type : String
#
    collection: 'invoices'
}
