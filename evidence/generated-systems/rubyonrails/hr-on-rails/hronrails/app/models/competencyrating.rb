class CompetencyRating < ApplicationRecord
  enum Rating: [:Unsatisfactory, :NeedsImprovement, :MeetsExpectations, :ExceedsExpectations, :Outstanding]


  has_many :Review, class_name: 'PerformanceReview'
  has_many :Competency, class_name: 'Competency'

end
