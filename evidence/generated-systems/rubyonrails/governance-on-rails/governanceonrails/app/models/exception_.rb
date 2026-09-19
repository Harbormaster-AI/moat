class Exception_ < ApplicationRecord
  enum ExceptionType: [:PolicyException, :ControlException, :RetentionException, :RiskAcceptance, :ComplianceWaiver]
  enum Status: [:Draft, :Submitted, :Approved, :Rejected, :Expired]


  has_many :RetentionSchedule, class_name: 'RetentionSchedule'
  has_many :Policy, class_name: 'Policy'
  has_many :Control, class_name: 'Control'
  has_many :Risk, class_name: 'Risk'

end
