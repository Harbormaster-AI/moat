class ReplenishmentPolicy < ApplicationRecord
  enum PolicyType: [:MinMax, :ReorderPoint, :EOQ, :Kanban]


  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Location, class_name: 'StorageLocation'

end
