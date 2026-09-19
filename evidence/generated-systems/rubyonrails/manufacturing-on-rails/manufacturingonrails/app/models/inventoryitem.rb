class InventoryItem < ApplicationRecord


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

  composed_of :serialId,
    class_name: "SerialId",
    mapping: [
      %w[serialId_value value]
    ]

  has_many :Item, class_name: 'Item'
  has_many :Location, class_name: 'Location'

end
