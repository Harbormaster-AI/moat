
// Define collection and schema for BillingAccount
export interface BillingAccount {
    accountNumber:
	type : string
    balance:
	type : Money
    Customer:
	type : Schema.Types.ObjectId
    Policies:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Policy' }]
    Invoices:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Invoice' }]
    Payments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Payment' }]
    Status:
 	type : String
#
    collection: 'billingAccounts'
}
