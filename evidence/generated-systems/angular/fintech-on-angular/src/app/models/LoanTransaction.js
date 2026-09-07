
// Define collection and schema for LoanTransaction
export interface LoanTransaction {
    transactionId:
	type : TransactionId
    amount:
	type : Money
    postingDate:
	type : Date
    Loan:
	type : Schema.Types.ObjectId
    Type:
 	type : String
    Status:
 	type : String
#
    collection: 'loanTransactions'
}
