class JobFamily < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :JobProfiles, class_name: 'JobProfile'

end
