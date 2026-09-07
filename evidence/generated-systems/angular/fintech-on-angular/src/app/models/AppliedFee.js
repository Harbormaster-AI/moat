
// Define collection and schema for AppliedFee
export interface AppliedFee {
    amount:
	type : Money
    description:
	type : string
    PaymentOrder:
	type : Schema.Types.ObjectId
    Transaction:
	type : Schema.Types.ObjectId
    FeeType:
 	type : String
#
    collection: 'appliedFees'
}
