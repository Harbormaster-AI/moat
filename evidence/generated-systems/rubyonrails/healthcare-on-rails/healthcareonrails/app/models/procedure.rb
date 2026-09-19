class Procedure < ApplicationRecord
  enum Status: [:Planned, :InProgress, :Completed, :Aborted]


  has_many :Encounter, class_name: 'Encounter'
  has_many :Performer, class_name: 'Clinician'
  has_many :ProcedureOrder, class_name: 'ProcedureOrder'

end
