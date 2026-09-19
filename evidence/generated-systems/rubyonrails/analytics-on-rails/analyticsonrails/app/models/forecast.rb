class Forecast < ApplicationRecord
  enum Granularity: [:Minute, :Hour, :Day, :Week, :Month, :Quarter, :Year]


  has_many :ModelVersion, class_name: 'ModelVersion'
  has_many :TimeSeries, class_name: 'TimeSeries'
  has_many :Datasets, class_name: 'DataSet'

end
