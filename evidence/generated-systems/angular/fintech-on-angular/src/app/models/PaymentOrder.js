
// Define collection and schema for PaymentOrder
export interface PaymentOrder {
    orderReference:
	type : string
    requestedExecutionDate:
	type : Date
    purpose:
	type : string
    SourceAccount:
	type : Schema.Types.ObjectId
    DestinationAccount:
	type : Schema.Types.ObjectId
    Beneficiary:
	type : Schema.Types.ObjectId
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Transaction' }]
    FxDeal:
	type : Schema.Types.ObjectId
    Fees:
 	type : [{ type: Schema.Types.ObjectId, ref: 'AppliedFee' }]
    PaymentMethod:
 	type : String
    Status:
 	type : String
    Priority:
 	type : String
#
    collection: 'paymentOrders'
}
