class Allergy < ApplicationRecord
  enum Severity: [:Mild, :Moderate, :Severe, :LifeThreatening]
  enum Status: [:Active, :Inactive, :Resolved, :EnteredInError]


  has_many :Patient, class_name: 'Patient'

end
