class MedicalSupplier < ApplicationRecord
  enum SupplierTier: [:Primary, :Secondary, :Distributor]


  has_many :Facilities, class_name: 'Facility'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
