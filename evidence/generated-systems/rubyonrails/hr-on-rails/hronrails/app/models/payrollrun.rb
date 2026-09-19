class PayrollRun < ApplicationRecord
  enum Status: [:Scheduled, :InProgress, :Completed, :Reversed]


  has_many :PayrollCalendar, class_name: 'PayrollCalendar'
  has_many :PayrollItems, class_name: 'PayrollItem'

end
