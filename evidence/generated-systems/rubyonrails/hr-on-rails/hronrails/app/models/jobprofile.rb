class JobProfile < ApplicationRecord
  enum JobLevel: [:Entry, :Intermediate, :Senior, :Lead, :Manager, :Director, :Executive]
  enum ExemptStatus: [:Exempt, :NonExempt]


  has_many :JobFamily, class_name: 'JobFamily'
  has_many :Competencies, class_name: 'Competency'
  has_many :TrainingRecommendations, class_name: 'TrainingCourse'
  has_many :Positions, class_name: 'Position'

end
