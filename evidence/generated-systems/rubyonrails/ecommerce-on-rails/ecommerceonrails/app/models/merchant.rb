class Merchant < ApplicationRecord


  has_many :Channels, class_name: 'Channel'
  has_many :Brands, class_name: 'Brand'
  has_many :FulfillmentCenters, class_name: 'FulfillmentCenter'
  has_many :TaxRules, class_name: 'TaxRule'
  has_many :PaymentProviders, class_name: 'PaymentProvider'
  has_many :Sellers, class_name: 'Seller'
  has_many :Promotions, class_name: 'Promotion'

end
