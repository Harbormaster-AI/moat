class Approval < ApplicationRecord
  enum Status: [:Pending, :Approved, :Rejected, :Cancelled]


  has_many :Approver, class_name: 'Employee'
  has_many :Timesheet, class_name: 'Timesheet'
  has_many :LeaveRequest, class_name: 'LeaveRequest'

end
