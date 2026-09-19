class Promotion < ApplicationRecord
  enum PromotionType: [:Catalog, :Cart, :Shipping]
  enum DiscountType: [:AmountOff, :PercentOff, :BuyXGetY, :FreeShipping]


  has_many :Merchant, class_name: 'Merchant'
  has_many :Channels, class_name: 'Channel'
  has_many :ApplicableProducts, class_name: 'Product'
  has_many :ApplicableCategories, class_name: 'Category'
  has_many :Coupons, class_name: 'Coupon'

end
