class Channel < ApplicationRecord
  enum ChannelType: [:Web, :MobileApp, :Marketplace, :Social, :POS]


  has_many :Merchant, class_name: 'Merchant'
  has_many :Catalogs, class_name: 'Catalog'
  has_many :Promotions, class_name: 'Promotion'
  has_many :ShippingMethods, class_name: 'ShippingMethod'
  has_many :PaymentProviders, class_name: 'PaymentProvider'

end
