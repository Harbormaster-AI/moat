class Observation < ApplicationRecord
  enum Interpretation: [:Normal, :AbnormalLow, :AbnormalHigh, :CriticalLow, :CriticalHigh, :Reactive, :Nonreactive, :Positive, :Negative]


  has_many :Encounter, class_name: 'Encounter'
  has_many :Patient, class_name: 'Patient'
  has_many :Device, class_name: 'MedicalDevice'
  has_many :LabResult, class_name: 'LabResult'

end
