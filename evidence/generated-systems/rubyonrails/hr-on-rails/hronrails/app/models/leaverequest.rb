class LeaveRequest < ApplicationRecord
  enum Status: [:Draft, :Submitted, :Approved, :Rejected, :Cancelled, :Taken]


  has_many :Employee, class_name: 'Employee'
  has_many :LeavePolicy, class_name: 'LeavePolicy'
  has_many :Approvals, class_name: 'Approval'

end
