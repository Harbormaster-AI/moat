class ComplianceRequirement < ApplicationRecord
  enum Applicability: [:Mandatory, :Recommended, :NotApplicable]
  enum Status: [:NotStarted, :InProgress, :Compliant, :NonCompliant, :Waived]


  has_many :ComplianceProgram, class_name: 'ComplianceProgram'
  has_many :Policies, class_name: 'Policy'
  has_many :Controls, class_name: 'Control'
  has_many :Obligations, class_name: 'Obligation'

end
