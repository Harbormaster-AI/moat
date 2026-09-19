class GoodsReceipt < ApplicationRecord
  enum Status: [:Open, :PartiallyProcessed, :Completed, :Rejected]


  has_many :PurchaseOrder, class_name: 'PurchaseOrder'
  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Lines, class_name: 'GoodsReceiptLine'

end
