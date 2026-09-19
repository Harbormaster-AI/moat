class CareTeam < ApplicationRecord
  enum CareSetting: [:Inpatient, :Outpatient, :Emergency, :HomeHealth, :Telehealth]


  has_many :Department, class_name: 'Department'
  has_many :Clinicians, class_name: 'Clinician'
  has_many :Patients, class_name: 'Patient'

end
