class EmploymentAssignment < ApplicationRecord
  enum AssignmentType: [:Primary, :Secondary, :Temporary]
  enum Status: [:Planned, :Active, :Completed, :Cancelled]


  has_many :Employee, class_name: 'Employee'
  has_many :Position, class_name: 'Position'
  has_many :Supervisor, class_name: 'Employee'

end
