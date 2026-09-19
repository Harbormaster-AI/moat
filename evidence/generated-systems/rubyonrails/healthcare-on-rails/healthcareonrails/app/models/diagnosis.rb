class Diagnosis < ApplicationRecord
  enum Certainty: [:Suspected, :Presumptive, :Confirmed, :RuledOut]


  has_many :Encounter, class_name: 'Encounter'
  has_many :Patient, class_name: 'Patient'

end
