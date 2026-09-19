class Employee < ApplicationRecord
  enum Status: [:Active, :OnLeave, :Suspended, :Terminated]


  composed_of :personName,
    class_name: "PersonName",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[personName_preferredName preferredName]
    ]

  composed_of :email,
    class_name: "Email",
    mapping: [
      %w[email_value value]
    ]

  composed_of :phoneNumber,
    class_name: "PhoneNumber",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[phoneNumber_extension extension]
    ]

  composed_of :nationalID,
    class_name: "NationalID",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[nationalID_type type]
    ]

  has_many :Manager, class_name: 'Employee'
  has_many :DirectReports, class_name: 'Employee'
  has_many :Department, class_name: 'Department'
  has_many :PrimaryLocation, class_name: 'Location'
  has_many :CostCenter, class_name: 'CostCenter'
  has_many :EmploymentAssignments, class_name: 'EmploymentAssignment'
  has_many :Contracts, class_name: 'EmploymentContract'
  has_many :BenefitEnrollments, class_name: 'BenefitEnrollment'
  has_many :Timesheets, class_name: 'Timesheet'
  has_many :LeaveRequests, class_name: 'LeaveRequest'
  has_many :PerformanceReviews, class_name: 'PerformanceReview'
  has_many :TrainingEnrollments, class_name: 'TrainingEnrollment'
  has_many :WorkAuthorizations, class_name: 'WorkAuthorization'

end
