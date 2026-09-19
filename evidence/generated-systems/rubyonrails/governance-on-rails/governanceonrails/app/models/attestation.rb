class Attestation < ApplicationRecord
  enum Result: [:Affirmative, :Negative, :Qualified]


  has_many :Control, class_name: 'Control'
  has_many :Policy, class_name: 'Policy'
  has_many :ComplianceProgram, class_name: 'ComplianceProgram'

end
