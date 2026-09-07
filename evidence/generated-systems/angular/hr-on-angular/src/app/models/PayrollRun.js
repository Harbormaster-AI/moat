
// Define collection and schema for PayrollRun
export interface PayrollRun {
    runNumber:
	type : string
    periodStart:
	type : Date
    periodEnd:
	type : Date
    paymentDate:
	type : Date
    PayrollCalendar:
	type : Schema.Types.ObjectId
    PayrollItems:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PayrollItem' }]
    Status:
 	type : String
#
    collection: 'payrollRuns'
}
