class AuditProgram < ApplicationRecord
  enum Cycle: [:Annual, :SemiAnnual, :Quarterly, :Continuous, :OneTime]
  enum Status: [:Planned, :Fieldwork, :Reporting, :Closed, :OnHold]


  has_many :Organization, class_name: 'Organization'
  has_many :Engagements, class_name: 'AuditEngagement'

end
