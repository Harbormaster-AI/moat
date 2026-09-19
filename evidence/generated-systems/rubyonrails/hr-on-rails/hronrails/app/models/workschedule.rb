class WorkSchedule < ApplicationRecord
  enum ScheduleType: [:Fixed, :Flexible, :Rotating]


  has_many :Contracts, class_name: 'EmploymentContract'
  has_many :Shifts, class_name: 'WorkShift'
  has_many :Exceptions, class_name: 'ScheduleException'

end
