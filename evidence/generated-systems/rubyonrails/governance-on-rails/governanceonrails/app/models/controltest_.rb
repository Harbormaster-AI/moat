class ControlTest_ < ApplicationRecord
  enum TestType: [:DesignEffectiveness, :OperatingEffectiveness, :Walkthrough, :Reperformance, :Inquiry, :Observation, :Inspection, :DataAnalysis]
  enum Effectiveness: [:Effective, :PartiallyEffective, :Ineffective, :NotTested]
  enum Status: [:Planned, :InProgress, :Completed, :Blocked, :Cancelled]


  has_many :Control, class_name: 'Control'
  has_many :Evidence, class_name: 'Evidence'
  has_many :Engagement, class_name: 'AuditEngagement'

end
