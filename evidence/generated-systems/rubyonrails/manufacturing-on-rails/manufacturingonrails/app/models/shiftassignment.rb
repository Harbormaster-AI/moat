class ShiftAssignment < ApplicationRecord


  has_many :Shift, class_name: 'Shift'
  has_many :Employee, class_name: 'Employee'
  has_many :WorkCenter, class_name: 'WorkCenter'

end
