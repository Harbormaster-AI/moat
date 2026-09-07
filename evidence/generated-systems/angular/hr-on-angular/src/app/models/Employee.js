
// Define collection and schema for Employee
export interface Employee {
    employeeNumber:
	type : string
    name:
	type : PersonName
    workEmail:
	type : Email
    workPhone:
	type : PhoneNumber
    dateOfHire:
	type : Date
    nationalId:
	type : NationalID
    Manager:
	type : Schema.Types.ObjectId
    DirectReports:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Employee' }]
    Department:
	type : Schema.Types.ObjectId
    PrimaryLocation:
	type : Schema.Types.ObjectId
    CostCenter:
	type : Schema.Types.ObjectId
    EmploymentAssignments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmploymentAssignment' }]
    Contracts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'EmploymentContract' }]
    BenefitEnrollments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'BenefitEnrollment' }]
    Timesheets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Timesheet' }]
    LeaveRequests:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LeaveRequest' }]
    PerformanceReviews:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PerformanceReview' }]
    TrainingEnrollments:
 	type : [{ type: Schema.Types.ObjectId, ref: 'TrainingEnrollment' }]
    WorkAuthorizations:
 	type : [{ type: Schema.Types.ObjectId, ref: 'WorkAuthorization' }]
    Status:
 	type : String
#
    collection: 'employees'
}
