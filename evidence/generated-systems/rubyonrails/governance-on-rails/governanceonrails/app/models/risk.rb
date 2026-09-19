class Risk < ApplicationRecord
  enum Category: [:Strategic, :Operational, :Financial, :Compliance, :Reputational, :Privacy, :Cybersecurity, :ThirdParty]
  enum Impact: [:Insignificant, :Minor, :Moderate, :Major, :Severe]
  enum Likelihood: [:Rare, :Unlikely, :Possible, :Likely, :AlmostCertain]
  enum Status: [:Identified, :Assessed, :Mitigated, :Accepted, :Transferred, :Closed]


  has_many :Organization, class_name: 'Organization'
  has_many :Controls, class_name: 'Control'
  has_many :Assessments, class_name: 'RiskAssessment'
  has_many :Issues, class_name: 'Issue'
  has_many :Findings, class_name: 'AuditFinding'

end
