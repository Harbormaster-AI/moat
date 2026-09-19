class InventoryThresholdAlert < ApplicationRecord
  enum AlertType: [:BelowMin, :AboveMax, :StockoutRisk, :ExcessStock, :ExpiryRisk]
  enum Status: [:New, :Acknowledged, :Resolved, :Dismissed]


  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Location, class_name: 'StorageLocation'
  has_many :RelatedPolicy, class_name: 'ReplenishmentPolicy'

end
