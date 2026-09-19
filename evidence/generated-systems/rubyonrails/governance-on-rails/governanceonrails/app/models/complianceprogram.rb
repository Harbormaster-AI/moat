class ComplianceProgram < ApplicationRecord
  enum Status: [:NotStarted, :InProgress, :Compliant, :NonCompliant, :Waived]


  has_many :Organization, class_name: 'Organization'
  has_many :Requirements, class_name: 'ComplianceRequirement'
  has_many :Controls, class_name: 'Control'
  has_many :Attestations, class_name: 'Attestation'
  has_many :Regulations, class_name: 'Regulation'

end
