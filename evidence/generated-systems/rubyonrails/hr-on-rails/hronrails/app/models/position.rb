class Position < ApplicationRecord
  enum Status: [:Open, :Filled, :Frozen, :Closed]
  enum WorkLocationType: [:Onsite, :Hybrid, :Remote]


  has_many :Department, class_name: 'Department'
  has_many :JobProfile, class_name: 'JobProfile'
  has_many :CostCenter, class_name: 'CostCenter'
  has_many :Location, class_name: 'Location'
  has_many :ManagerPosition, class_name: 'Position'
  has_many :DirectReports, class_name: 'Position'
  has_many :Assignments, class_name: 'EmploymentAssignment'

end
