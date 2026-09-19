class Supplier < ApplicationRecord
  enum Status: [:Active, :Inactive, :Onboarding, :Suspended]


  has_many :Merchant, class_name: 'Merchant'
  has_many :Products, class_name: 'Product'
  has_many :FulfillmentCenters, class_name: 'FulfillmentCenter'

end
