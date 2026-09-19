class ScheduleException < ApplicationRecord


  has_many :WorkSchedule, class_name: 'WorkSchedule'
  has_many :Employee, class_name: 'Employee'

end
