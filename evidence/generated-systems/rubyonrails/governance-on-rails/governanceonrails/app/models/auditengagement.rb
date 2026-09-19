class AuditEngagement < ApplicationRecord
  enum Status: [:Planned, :Fieldwork, :Reporting, :Closed, :OnHold]


  has_many :AuditProgram, class_name: 'AuditProgram'
  has_many :BusinessUnits, class_name: 'BusinessUnit'
  has_many :ControlTests, class_name: 'ControlTest_'
  has_many :Workpapers, class_name: 'AuditWorkpaper'
  has_many :Findings, class_name: 'AuditFinding'

end
