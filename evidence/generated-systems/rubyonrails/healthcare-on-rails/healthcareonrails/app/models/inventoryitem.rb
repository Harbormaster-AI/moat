class InventoryItem < ApplicationRecord


  has_many :Facility, class_name: 'Facility'
  has_many :Supplier, class_name: 'MedicalSupplier'

end
