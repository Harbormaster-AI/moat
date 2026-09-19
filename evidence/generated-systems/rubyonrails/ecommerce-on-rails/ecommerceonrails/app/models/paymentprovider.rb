class PaymentProvider < ApplicationRecord
  enum ProviderType: [:PSP, :Gateway, :Aggregator, :Manual]


  has_many :Merchant, class_name: 'Merchant'
  has_many :Channels, class_name: 'Channel'
  has_many :Payments, class_name: 'Payment'
  has_many :Subscriptions, class_name: 'Subscription'

end
