class Product < ApplicationRecord
  enum ProductType: [:Physical, :Digital, :Service, :Bundle, :Subscription]
  enum DefaultTaxClass: [:Standard, :Reduced, :Zero, :Exempt, :DigitalServices, :Food, :Clothing]


  has_many :Brand, class_name: 'Brand'
  has_many :Categories, class_name: 'Category'
  has_many :Variants, class_name: 'ProductVariant'
  has_many :MediaAssets, class_name: 'MediaAsset'
  has_many :Reviews, class_name: 'Review'
  has_many :Seller, class_name: 'Seller'

end
