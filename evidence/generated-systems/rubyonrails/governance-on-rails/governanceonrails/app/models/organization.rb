class Organization < ApplicationRecord


  has_many :GovernanceBodies, class_name: 'GovernanceBody'
  has_many :Policies, class_name: 'Policy'
  has_many :Risks, class_name: 'Risk'
  has_many :ThirdParties, class_name: 'ThirdParty'
  has_many :RecordsRepositories, class_name: 'RecordsRepository'
  has_many :DataProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :CompliancePrograms, class_name: 'ComplianceProgram'
  has_many :AuditPrograms, class_name: 'AuditProgram'
  has_many :BusinessUnits, class_name: 'BusinessUnit'
  has_many :Matters, class_name: 'Matter'
  has_many :DataBreaches, class_name: 'DataBreach'

end
