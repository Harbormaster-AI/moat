class ProductVariant < ApplicationRecord
  enum WeightUnit: [:Gram, :Kilogram, :Ounce, :Pound]


  composed_of :sKU,
    class_name: "SKU",
    mapping: [
      %w[sKU_value value]
    ]

  has_many :Product, class_name: 'Product'
  has_many :Pricing, class_name: 'ProductPricing'
  has_many :InventoryItems, class_name: 'InventoryItem'
  has_many :MediaAssets, class_name: 'MediaAsset'
  has_many :Subscriptions, class_name: 'Subscription'
  has_many :CartItems, class_name: 'CartItem'
  has_many :OrderLines, class_name: 'OrderLine'
  has_many :WishlistItems, class_name: 'WishlistItem'

end
