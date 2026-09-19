class PerformanceReview < ApplicationRecord
  enum Rating: [:Unsatisfactory, :NeedsImprovement, :MeetsExpectations, :ExceedsExpectations, :Outstanding]
  enum Status: [:NotStarted, :InProgress, :Finalized, :Acknowledged]


  has_many :Employee, class_name: 'Employee'
  has_many :Reviewer, class_name: 'Employee'
  has_many :Cycle, class_name: 'PerformanceCycle'
  has_many :CompetencyRatings, class_name: 'CompetencyRating'
  has_many :Goals, class_name: 'Goal'

end
