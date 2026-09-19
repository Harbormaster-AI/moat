class Appointment < ApplicationRecord
  enum Status: [:Proposed, :Booked, :Arrived, :Fulfilled, :Cancelled, :NoShow, :EnteredInError]
  enum Priority: [:Routine, :Urgent, :Stat]


  has_many :Patient, class_name: 'Patient'
  has_many :Clinician, class_name: 'Clinician'
  has_many :Facility, class_name: 'Facility'
  has_many :Encounter, class_name: 'Encounter'

end
