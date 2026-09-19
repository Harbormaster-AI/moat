class TrainingCourse < ApplicationRecord
  enum DeliveryMethod: [:Classroom, :Virtual, :SelfPaced, :Blended]


  has_many :Prerequisites, class_name: 'TrainingCourse'
  has_many :Enrollments, class_name: 'TrainingEnrollment'
  has_many :JobProfiles, class_name: 'JobProfile'

end
