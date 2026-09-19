class PayrollCalendar < ApplicationRecord
  enum PayFrequency: [:Weekly, :Biweekly, :Semimonthly, :Monthly, :Quarterly]


  has_many :Organization, class_name: 'Organization'
  has_many :PayrollRuns, class_name: 'PayrollRun'
  has_many :Employees, class_name: 'Employee'

end
