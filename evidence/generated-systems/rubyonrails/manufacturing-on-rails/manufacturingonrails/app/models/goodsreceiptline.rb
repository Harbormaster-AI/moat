class GoodsReceiptLine < ApplicationRecord


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  composed_of :lotId,
    class_name: "LotId",
    mapping: [
      %w[lotId_value value]
    ]

  has_many :GoodsReceipt, class_name: 'GoodsReceipt'
  has_many :Item, class_name: 'Item'
  has_many :InventoryTransaction, class_name: 'InventoryTransaction'

end
