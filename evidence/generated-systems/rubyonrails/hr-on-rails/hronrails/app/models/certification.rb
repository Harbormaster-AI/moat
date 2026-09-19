class Certification < ApplicationRecord


  has_many :Employee, class_name: 'Employee'
  has_many :Course, class_name: 'TrainingCourse'

end
