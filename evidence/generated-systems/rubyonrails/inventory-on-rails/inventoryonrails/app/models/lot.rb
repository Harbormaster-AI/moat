class Lot < ApplicationRecord
  enum LotStatus: [:Released, :Quarantined, :Expired, :Blocked, :PendingTest]


  composed_of :batchNumber,
    class_name: "BatchNumber",
    mapping: [
      %w[batchNumber_value value]
    ]

  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
