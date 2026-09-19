class DataPipeline < ApplicationRecord
  enum TriggerType: [:Manual, :Schedule, :Event]
  enum Status: [:Draft, :Active, :Paused, :Failed, :Succeeded]


  composed_of :cronSchedule,
    class_name: "CronSchedule",
    mapping: [
      ${$mapping}, 
      %w[cronSchedule_timezone timezone]
    ]

  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :Tasks, class_name: 'DataTask'
  has_many :Sources, class_name: 'DataSource'
  has_many :Outputs, class_name: 'DataSet'
  has_many :LineageNode, class_name: 'LineageNode'

end
