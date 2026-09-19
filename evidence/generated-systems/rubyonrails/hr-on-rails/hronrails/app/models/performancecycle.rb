class PerformanceCycle < ApplicationRecord
  enum Status: [:Planned, :Open, :Closed]


  has_many :Organization, class_name: 'Organization'
  has_many :Reviews, class_name: 'PerformanceReview'
  has_many :Goals, class_name: 'Goal'

end
