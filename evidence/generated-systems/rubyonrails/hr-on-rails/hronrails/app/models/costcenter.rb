class CostCenter < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :Departments, class_name: 'Department'
  has_many :Positions, class_name: 'Position'
  has_many :Employees, class_name: 'Employee'

end
