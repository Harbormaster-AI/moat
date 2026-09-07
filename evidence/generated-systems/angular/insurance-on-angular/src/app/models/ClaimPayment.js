
// Define collection and schema for ClaimPayment
export interface ClaimPayment {
    paymentNumber:
	type : string
    amount:
	type : Money
    paymentDate:
	type : Date
    Claim:
	type : Schema.Types.ObjectId
    Exposure:
	type : Schema.Types.ObjectId
    Beneficiary:
	type : Schema.Types.ObjectId
    ServiceProvider:
	type : Schema.Types.ObjectId
    Customer:
	type : Schema.Types.ObjectId
    PayeeType:
 	type : String
    Method:
 	type : String
    Status:
 	type : String
#
    collection: 'claimPayments'
}
