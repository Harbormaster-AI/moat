class CarePlan < ApplicationRecord
  enum Status: [:Draft, :Active, :Suspended, :Completed, :Cancelled]


  has_many :Patient, class_name: 'Patient'
  has_many :Encounters, class_name: 'Encounter'
  has_many :Tasks, class_name: 'CareTask'
  has_many :CareTeam, class_name: 'CareTeam'

end
