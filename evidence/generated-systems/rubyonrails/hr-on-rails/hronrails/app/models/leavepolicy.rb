class LeavePolicy < ApplicationRecord
  enum LeaveCategory: [:Vacation, :Sick, :Parental, :Bereavement, :Unpaid, :JuryDuty]
  enum AccrualUnit: [:Hours, :Days]


  has_many :Organization, class_name: 'Organization'
  has_many :LeaveRequests, class_name: 'LeaveRequest'

end
