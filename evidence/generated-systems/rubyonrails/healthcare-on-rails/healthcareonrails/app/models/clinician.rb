class Clinician < ApplicationRecord
  enum ClinicianType: [:Physician, :NursePractitioner, :PhysicianAssistant, :RegisteredNurse, :Pharmacist, :Therapist, :Technician]
  enum Specialty: [:InternalMedicine, :FamilyMedicine, :Cardiology, :Oncology, :Orthopedics, :Pediatrics, :Radiology, :Pathology, :Anesthesiology, :Surgery, :Psychiatry]


  has_many :CareTeams, class_name: 'CareTeam'
  has_many :Appointments, class_name: 'Appointment'
  has_many :Encounters, class_name: 'Encounter'
  has_many :Procedures, class_name: 'Procedure'
  has_many :ImagingReports, class_name: 'ImagingReport'

end
