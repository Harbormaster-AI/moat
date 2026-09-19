class PurchaseOrder < ApplicationRecord
  enum Status: [:Draft, :Submitted, :Acknowledged, :PartiallyReceived, :Received, :Closed, :Cancelled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Supplier, class_name: 'Supplier'
  has_many :Plant, class_name: 'Plant'
  has_many :Lines, class_name: 'PurchaseOrderLine'
  has_many :GoodsReceipts, class_name: 'GoodsReceipt'

end
