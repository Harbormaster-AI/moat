class TrainingEnrollment < ApplicationRecord
  enum Status: [:Enrolled, :InProgress, :Completed, :Failed, :Cancelled]


  has_many :Course, class_name: 'TrainingCourse'
  has_many :Employee, class_name: 'Employee'
  has_many :Instructor, class_name: 'Employee'

end
