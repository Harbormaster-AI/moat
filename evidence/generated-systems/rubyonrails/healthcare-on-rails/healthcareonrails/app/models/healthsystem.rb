class HealthSystem < ApplicationRecord


  has_many :Facilities, class_name: 'Facility'
  has_many :Suppliers, class_name: 'MedicalSupplier'

end
