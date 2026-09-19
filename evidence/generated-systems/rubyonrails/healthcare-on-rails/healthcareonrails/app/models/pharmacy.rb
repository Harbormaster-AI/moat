class Pharmacy < ApplicationRecord


  has_many :Facility, class_name: 'Facility'
  has_many :MedicationDispenses, class_name: 'MedicationDispense'
  has_many :MedicationOrders, class_name: 'MedicationOrder'

end
