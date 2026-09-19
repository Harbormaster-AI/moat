class Competency < ApplicationRecord


  has_many :JobProfiles, class_name: 'JobProfile'
  has_many :CompetencyRatings, class_name: 'CompetencyRating'

end
