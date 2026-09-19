class EmploymentContract < ApplicationRecord
  enum EmploymentType: [:FullTime, :PartTime, :Temporary, :Intern, :Contractor, :Seasonal]
  enum Status: [:Draft, :Active, :Suspended, :Expired, :Terminated]
  enum PayFrequency: [:Weekly, :Biweekly, :Semimonthly, :Monthly, :Quarterly]


  has_many :Employee, class_name: 'Employee'
  has_many :CompensationPackage, class_name: 'CompensationPackage'
  has_many :WorkSchedule, class_name: 'WorkSchedule'
  has_many :Location, class_name: 'Location'
  has_many :PayrollCalendar, class_name: 'PayrollCalendar'

end
