class Timesheet < ApplicationRecord
  enum Status: [:Draft, :Submitted, :Approved, :Rejected, :Processed]


  has_many :Employee, class_name: 'Employee'
  has_many :TimeEntries, class_name: 'TimeEntry'
  has_many :Approvals, class_name: 'Approval'

end
