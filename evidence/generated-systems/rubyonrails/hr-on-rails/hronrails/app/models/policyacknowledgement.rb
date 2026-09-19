class PolicyAcknowledgement < ApplicationRecord
  enum Status: [:Pending, :Acknowledged, :Declined]


  has_many :Policy, class_name: 'Policy'
  has_many :Employee, class_name: 'Employee'

end
