class Seller < ApplicationRecord
  enum Status: [:Active, :Inactive, :Suspended]


  has_many :Merchant, class_name: 'Merchant'
  has_many :Products, class_name: 'Product'
  has_many :Payouts, class_name: 'Payout'
  has_many :Orders, class_name: 'Order'

end
