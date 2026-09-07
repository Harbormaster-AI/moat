
// Define collection and schema for Payment
export interface Payment {
    paymentNumber:
	type : string
    amount:
	type : Money
    paymentDate:
	type : Date
    Invoice:
	type : Schema.Types.ObjectId
    Payer:
	type : Schema.Types.ObjectId
    Method:
 	type : String
#
    collection: 'payments'
}
