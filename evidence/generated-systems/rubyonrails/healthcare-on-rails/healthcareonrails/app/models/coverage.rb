class Coverage < ApplicationRecord
  enum CoverageType: [:Medical, :Pharmacy, :Dental, :Vision, :BehavioralHealth]


  has_many :Patient, class_name: 'Patient'
  has_many :Plan, class_name: 'InsurancePlan'
  has_many :Claims, class_name: 'Claim'
  has_many :Authorizations, class_name: 'Authorization'

end
