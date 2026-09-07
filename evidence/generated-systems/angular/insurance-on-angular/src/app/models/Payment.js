
// Define collection and schema for Payment
export interface Payment {
    paymentReference:
	type : string
    amount:
	type : Money
    paymentDate:
	type : Date
    Invoice:
	type : Schema.Types.ObjectId
    BillingAccount:
	type : Schema.Types.ObjectId
    Policy:
	type : Schema.Types.ObjectId
    Method:
 	type : String
    Status:
 	type : String
#
    collection: 'payments'
}
