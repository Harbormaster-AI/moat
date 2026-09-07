
// Define collection and schema for PaymentContract
export interface PaymentContract {
    contractNumber:
	type : string
    pricingPlanCode:
	type : string
    Merchant:
	type : Schema.Types.ObjectId
    Acquirer:
	type : Schema.Types.ObjectId
    Status:
 	type : String
#
    collection: 'paymentContracts'
}
