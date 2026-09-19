class ClinicalOrder < ApplicationRecord
  enum Status: [:Draft, :Active, :OnHold, :Completed, :Cancelled]
  enum OrderType: [:Medication, :Laboratory, :Imaging, :Procedure, :Consultation]
  enum Priority: [:Routine, :Urgent, :Stat]


  has_many :Patient, class_name: 'Patient'
  has_many :Encounter, class_name: 'Encounter'
  has_many :OrderingClinician, class_name: 'Clinician'
  has_many :MedicationOrders, class_name: 'MedicationOrder'
  has_many :LaboratoryOrders, class_name: 'LaboratoryOrder'
  has_many :ImagingOrders, class_name: 'ImagingOrder'
  has_many :ProcedureOrders, class_name: 'ProcedureOrder'
  has_many :Authorizations, class_name: 'Authorization'

end
