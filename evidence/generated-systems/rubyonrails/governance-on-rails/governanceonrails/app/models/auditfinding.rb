class AuditFinding < ApplicationRecord
  enum Severity: [:Low, :Medium, :High, :Critical]
  enum Status: [:Open, :InRemediation, :Validated, :Closed]


  has_many :Engagement, class_name: 'AuditEngagement'
  has_many :Workpaper, class_name: 'AuditWorkpaper'
  has_many :CorrectiveActions, class_name: 'CorrectiveAction'
  has_many :RelatedRisks, class_name: 'Risk'
  has_many :RelatedControls, class_name: 'Control'
  has_many :Issues, class_name: 'Issue'

end
