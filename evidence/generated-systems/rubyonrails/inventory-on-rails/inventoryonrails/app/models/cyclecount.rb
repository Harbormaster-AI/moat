class CycleCount < ApplicationRecord
  enum Status: [:Planned, :InProgress, :Completed, :Posted, :Cancelled]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Locations, class_name: 'StorageLocation'
  has_many :Entries, class_name: 'CycleCountEntry'
  has_many :Transactions, class_name: 'InventoryTransaction'

end
