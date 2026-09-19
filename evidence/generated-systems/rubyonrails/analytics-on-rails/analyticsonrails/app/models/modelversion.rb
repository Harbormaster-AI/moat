class ModelVersion < ApplicationRecord
  enum Lifecycle: [:Draft, :Staging, :Production, :Archived]
  enum TrainingStatus: [:Queued, :Running, :Completed, :Failed]


  has_many :Model, class_name: 'Model'
  has_many :TrainingRun, class_name: 'TrainingRun'
  has_many :EvaluationMetrics, class_name: 'EvaluationMetric'
  has_many :Deployments, class_name: 'InferenceEndpoint'
  has_many :FeatureSets, class_name: 'FeatureSet'
  has_many :Datasets, class_name: 'DataSet'

end
