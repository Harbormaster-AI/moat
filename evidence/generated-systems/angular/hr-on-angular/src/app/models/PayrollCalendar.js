
// Define collection and schema for PayrollCalendar
export interface PayrollCalendar {
    name:
	type : string
    country:
	type : string
    Organization:
	type : Schema.Types.ObjectId
    PayrollRuns:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PayrollRun' }]
    Employees:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Employee' }]
    PayFrequency:
 	type : String
#
    collection: 'payrollCalendars'
}
