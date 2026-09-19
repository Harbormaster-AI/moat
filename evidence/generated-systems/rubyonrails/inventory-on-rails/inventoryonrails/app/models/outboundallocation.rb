class OutboundAllocation < ApplicationRecord
  enum Status: [:Proposed, :Confirmed, :Picked, :Short, :Cancelled]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :InventoryItem, class_name: 'InventoryItem'
  has_many :Reservation, class_name: 'Reservation'
  has_many :Lot, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'
  has_many :SourceLocation, class_name: 'StorageLocation'

end
