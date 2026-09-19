class BOMItem < ApplicationRecord


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Bom, class_name: 'BOM'
  has_many :Component, class_name: 'Item'

end
