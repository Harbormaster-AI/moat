class TimeSeries < ApplicationRecord
  enum Granularity: [:Minute, :Hour, :Day, :Week, :Month, :Quarter, :Year]


  has_many :Datasets, class_name: 'DataSet'
  has_many :Forecasts, class_name: 'Forecast'
  has_many :Anomalies, class_name: 'Anomaly'

end
