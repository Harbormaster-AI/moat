class Encounter < ApplicationRecord
  enum Status: [:Planned, :InProgress, :OnHold, :Discharged, :Cancelled]
  enum EncounterType: [:Inpatient, :Outpatient, :Emergency, :Observation, :Telemedicine]


  has_many :Patient, class_name: 'Patient'
  has_many :Clinician, class_name: 'Clinician'
  has_many :Facility, class_name: 'Facility'
  has_many :Appointment, class_name: 'Appointment'
  has_many :Diagnoses, class_name: 'Diagnosis'
  has_many :Procedures, class_name: 'Procedure'
  has_many :Observations, class_name: 'Observation'
  has_many :Orders, class_name: 'ClinicalOrder'
  has_many :Admission, class_name: 'Admission'
  has_many :Discharge, class_name: 'Discharge'

end
