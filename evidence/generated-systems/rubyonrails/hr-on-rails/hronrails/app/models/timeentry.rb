class TimeEntry < ApplicationRecord
  enum EntryType: [:Regular, :Overtime, :Sick, :Vacation, :Unpaid]


  has_many :Timesheet, class_name: 'Timesheet'
  has_many :Employee, class_name: 'Employee'
  has_many :CostCenter, class_name: 'CostCenter'

end
