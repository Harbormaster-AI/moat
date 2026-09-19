class Department < ApplicationRecord
  enum DepartmentType: [:Emergency, :Cardiology, :Oncology, :Orthopedics, :Pediatrics, :Radiology, :Pathology, :Pharmacy, :IntensiveCare]


  has_many :Facility, class_name: 'Facility'
  has_many :CareTeams, class_name: 'CareTeam'

end
