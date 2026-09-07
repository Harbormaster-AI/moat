
// Define collection and schema for PaymentMethod
export interface PaymentMethod {
    last4:
	type : string
    cardholderName:
	type : string
    billingAddress:
	type : Address
    BillingProfile:
	type : Schema.Types.ObjectId
    MethodType:
 	type : String
#
    collection: 'paymentMethods'
}
