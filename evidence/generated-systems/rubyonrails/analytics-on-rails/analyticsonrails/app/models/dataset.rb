class DataSet < ApplicationRecord
  enum DataFormat: [:CSV, :JSON, :Parquet, :Avro, :ORC, :XML]


  composed_of :cronSchedule,
    class_name: "CronSchedule",
    mapping: [
      ${$mapping}, 
      %w[cronSchedule_timezone timezone]
    ]

  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :Sources, class_name: 'DataSource'
  has_many :Pipelines, class_name: 'DataPipeline'
  has_many :SemanticModels, class_name: 'SemanticModel'
  has_many :Dimensions, class_name: 'Dimension'
  has_many :Measures, class_name: 'Measure'
  has_many :Metrics, class_name: 'Metric'
  has_many :QualityRules, class_name: 'QualityRule'
  has_many :LineageNode, class_name: 'LineageNode'
  has_many :Tags, class_name: 'Tag'

end
