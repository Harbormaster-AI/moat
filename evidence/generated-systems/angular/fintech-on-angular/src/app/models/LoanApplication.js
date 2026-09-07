
// Define collection and schema for LoanApplication
export interface LoanApplication {
    applicationNumber:
	type : string
    amountRequested:
	type : Money
    termMonths:
	type : number
    submittedAt:
	type : DateTime
    Customer:
	type : Schema.Types.ObjectId
    RiskAssessment:
	type : Schema.Types.ObjectId
    Loan:
	type : Schema.Types.ObjectId
    Product:
 	type : String
    Purpose:
 	type : String
    Status:
 	type : String
#
    collection: 'loanApplications'
}
