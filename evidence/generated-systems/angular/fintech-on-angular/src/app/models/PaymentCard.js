
// Define collection and schema for PaymentCard
export interface PaymentCard {
    cardToken:
	type : CardNumberToken
    maskedPan:
	type : string
    expiryMonth:
	type : number
    expiryYear:
	type : number
    cardholderName:
	type : string
    Customer:
	type : Schema.Types.ObjectId
    Account:
	type : Schema.Types.ObjectId
    Tokenizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'CardTokenization' }]
    Disputes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dispute' }]
    Scheme:
 	type : String
    Status:
 	type : String
#
    collection: 'paymentCards'
}
