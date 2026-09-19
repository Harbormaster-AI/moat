class Condition < ApplicationRecord
  enum ClinicalStatus: [:Active, :Recurrence, :Relapse, :Remission, :Resolved]
  enum VerificationStatus: [:Suspected, :Presumptive, :Confirmed, :RuledOut]


  has_many :Patient, class_name: 'Patient'

end
