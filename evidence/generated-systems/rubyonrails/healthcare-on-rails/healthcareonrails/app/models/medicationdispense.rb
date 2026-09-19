class MedicationDispense < ApplicationRecord
  enum Status: [:Preparation, :InProgress, :Completed, :Cancelled]


  has_many :MedicationOrder, class_name: 'MedicationOrder'
  has_many :Pharmacy, class_name: 'Pharmacy'
  has_many :Patient, class_name: 'Patient'

end
