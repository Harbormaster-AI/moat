class Department < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :Manager, class_name: 'Employee'
  has_many :Positions, class_name: 'Position'
  has_many :Employees, class_name: 'Employee'
  has_many :CostCenter, class_name: 'CostCenter'

end
