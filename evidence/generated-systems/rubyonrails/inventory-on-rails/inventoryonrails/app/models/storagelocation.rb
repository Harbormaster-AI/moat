class StorageLocation < ApplicationRecord
  enum LocationType: [:Bin, :Bulk, :Staging, :Dock, :Picking, :Packing, :Quality, :Return, :ColdStorage]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :ParentLocation, class_name: 'StorageLocation'
  has_many :ChildLocations, class_name: 'StorageLocation'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
