class StockAdjustment < ApplicationRecord
  enum AdjustmentType: [:Increase, :Decrease, :Reclassification]
  enum Status: [:Draft, :Approved, :Posted, :Cancelled]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Lines, class_name: 'StockAdjustmentLine'
  has_many :Transactions, class_name: 'InventoryTransaction'

end
