
// Define collection and schema for RepaymentSchedule
export interface RepaymentSchedule {
    installmentNumber:
	type : number
    dueDate:
	type : Date
    amountDue:
	type : Money
    principalDue:
	type : Money
    interestDue:
	type : Money
    Loan:
	type : Schema.Types.ObjectId
    Payments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Transaction' }]
    Status:
 	type : String
#
    collection: 'repaymentSchedules'
}
