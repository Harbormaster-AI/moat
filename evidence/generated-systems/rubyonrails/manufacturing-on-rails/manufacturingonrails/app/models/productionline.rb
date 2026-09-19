class ProductionLine < ApplicationRecord
  enum LineType: [:Discrete, :Batch, :Continuous, :FlexibleCell]


  has_many :Plant, class_name: 'Plant'
  has_many :WorkCenters, class_name: 'WorkCenter'

end
