class LineageNode < ApplicationRecord
  enum NodeType: [:Dataset, :Pipeline, :Model, :Dashboard, :Report, :FeatureSet, :Notebook]


  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :Inputs, class_name: 'LineageNode'
  has_many :Outputs, class_name: 'LineageNode'
  has_many :Datasets, class_name: 'DataSet'
  has_many :Models, class_name: 'Model'
  has_many :Pipelines, class_name: 'DataPipeline'
  has_many :Dashboards, class_name: 'Dashboard'
  has_many :Reports, class_name: 'Report'

end
