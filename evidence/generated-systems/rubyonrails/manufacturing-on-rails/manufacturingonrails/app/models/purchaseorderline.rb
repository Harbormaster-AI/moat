class PurchaseOrderLine < ApplicationRecord


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :PurchaseOrder, class_name: 'PurchaseOrder'
  has_many :Item, class_name: 'Item'

end
