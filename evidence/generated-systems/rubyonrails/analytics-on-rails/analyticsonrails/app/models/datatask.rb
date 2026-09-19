class DataTask < ApplicationRecord
  enum TaskType: [:Extract, :Transform, :Load, :Validate, :Enrich]


  has_many :Pipeline, class_name: 'DataPipeline'
  has_many :InputDatasets, class_name: 'DataSet'
  has_many :OutputDatasets, class_name: 'DataSet'

end
