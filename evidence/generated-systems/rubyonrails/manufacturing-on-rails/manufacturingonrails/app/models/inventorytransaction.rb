class InventoryTransaction < ApplicationRecord
  enum TransactionType: [:Receipt, :Issue, :Return, :Adjustment, :Transfer, :Consumption, :ProductionReceipt, :Scrap]


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  has_many :Item, class_name: 'Item'
  has_many :Location, class_name: 'Location'
  has_many :WorkOrder, class_name: 'WorkOrder'
  has_many :PurchaseOrder, class_name: 'PurchaseOrder'
  has_many :SalesOrder, class_name: 'SalesOrder'

end
