class Quarantine < ApplicationRecord
  enum Disposition: [:Release, :Scrap, :ReturnToVendor, :Rework]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Items, class_name: 'InventoryItem'
  has_many :Lot, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'

end
