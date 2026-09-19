class Issue < ApplicationRecord
  enum IssueType: [:ControlDeficiency, :ProcessGap, :ComplianceBreach, :SecurityIncident, :DataQualityIssue, :ThirdPartyIssue]
  enum Priority: [:Low, :Medium, :High, :Urgent]
  enum Status: [:Open, :Investigating, :RemediationPlanned, :RemediationInProgress, :Verified, :Closed]


  has_many :Risk, class_name: 'Risk'
  has_many :Finding, class_name: 'AuditFinding'
  has_many :CorrectiveActions, class_name: 'CorrectiveAction'
  has_many :Control, class_name: 'Control'

end
