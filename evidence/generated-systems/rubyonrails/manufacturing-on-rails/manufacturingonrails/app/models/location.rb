class Location < ApplicationRecord
  enum LocationType: [:Bin, :Dock, :Staging, :QAHold, :Scrap]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
