
// Define collection and schema for PaymentMethod
export interface PaymentMethod {
    preferred:
	type : boolean
    Employee:
	type : Schema.Types.ObjectId
    BankAccount:
	type : Schema.Types.ObjectId
    MethodType:
 	type : String
#
    collection: 'paymentMethods'
}
