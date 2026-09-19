class SerialNumber < ApplicationRecord
  enum Status: [:Active, :Assigned, :InTransit, :Consumed, :Returned, :Scrapped]


  composed_of :serialCode,
    class_name: "SerialCode",
    mapping: [
      %w[serialCode_value value]
    ]

  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :CurrentInventoryItem, class_name: 'InventoryItem'
  has_many :Lot, class_name: 'Lot'

end
