class InventoryItem < ApplicationRecord
  enum Status: [:InStock, :LowStock, :OutOfStock, :Backorder, :Preorder]


  has_many :Variant, class_name: 'ProductVariant'
  has_many :FulfillmentCenter, class_name: 'FulfillmentCenter'

end
