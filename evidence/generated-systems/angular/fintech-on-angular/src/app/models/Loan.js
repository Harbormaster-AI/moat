
// Define collection and schema for Loan
export interface Loan {
    loanNumber:
	type : string
    principal:
	type : Money
    interestRate:
	type : String
    originationDate:
	type : Date
    maturityDate:
	type : Date
    Customer:
	type : Schema.Types.ObjectId
    Schedule:
 	type : [{ type: Schema.Types.ObjectId, ref: 'RepaymentSchedule' }]
    Collateral:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Collateral' }]
    Transactions:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LoanTransaction' }]
    RateType:
 	type : String
    Status:
 	type : String
#
    collection: 'loans'
}
