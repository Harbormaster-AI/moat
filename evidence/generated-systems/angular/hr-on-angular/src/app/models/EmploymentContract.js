
// Define collection and schema for EmploymentContract
export interface EmploymentContract {
    contractNumber:
	type : string
    startDate:
	type : Date
    endDate:
	type : Date
    workHoursPerWeek:
	type : String
    Employee:
	type : Schema.Types.ObjectId
    CompensationPackage:
	type : Schema.Types.ObjectId
    WorkSchedule:
	type : Schema.Types.ObjectId
    Location:
	type : Schema.Types.ObjectId
    PayrollCalendar:
	type : Schema.Types.ObjectId
    EmploymentType:
 	type : String
    Status:
 	type : String
    PayFrequency:
 	type : String
#
    collection: 'employmentContracts'
}
