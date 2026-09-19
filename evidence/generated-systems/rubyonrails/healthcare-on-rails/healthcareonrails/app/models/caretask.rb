class CareTask < ApplicationRecord
  enum Status: [:Requested, :Accepted, :InProgress, :Completed, :Cancelled, :Failed]
  enum Priority: [:Routine, :Urgent, :Stat]


  has_many :CarePlan, class_name: 'CarePlan'
  has_many :AssignedTo, class_name: 'Clinician'
  has_many :Encounter, class_name: 'Encounter'

end
