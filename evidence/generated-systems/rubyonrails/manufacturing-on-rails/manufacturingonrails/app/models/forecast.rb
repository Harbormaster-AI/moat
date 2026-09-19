class Forecast < ApplicationRecord
  enum Method: [:MovingAverage, :ExponentialSmoothing, :Croston, :ARIMA, :Manual]


  has_many :Lines, class_name: 'ForecastLine'

end
