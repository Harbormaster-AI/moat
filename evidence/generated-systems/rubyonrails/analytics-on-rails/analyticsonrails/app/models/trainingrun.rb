class TrainingRun < ApplicationRecord
  enum Status: [:Queued, :Running, :Completed, :Failed]


  has_many :Experiment, class_name: 'Experiment'
  has_many :ModelVersion, class_name: 'ModelVersion'
  has_many :InputDatasets, class_name: 'DataSet'
  has_many :Features, class_name: 'Feature'
  has_many :RunMetrics, class_name: 'RunMetric'
  has_many :RunParameters, class_name: 'RunParameter'

end
