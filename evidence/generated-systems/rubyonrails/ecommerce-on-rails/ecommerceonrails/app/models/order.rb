class Order < ApplicationRecord
  enum Status: [:Pending, :Confirmed, :Paid, :PartiallyShipped, :Shipped, :Delivered, :Cancelled, :Refunded, :PartiallyRefunded]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Channel, class_name: 'Channel'
  has_many :OrderLines, class_name: 'OrderLine'
  has_many :Payments, class_name: 'Payment'
  has_many :Shipments, class_name: 'Shipment'
  has_many :Refunds, class_name: 'Refund'
  has_many :AppliedPromotions, class_name: 'Promotion'
  has_many :Seller, class_name: 'Seller'
  has_many :GiftCardRedemptions, class_name: 'GiftCardRedemption'
  has_many :CouponRedemptions, class_name: 'CouponRedemption'
  has_many :ReturnRequests, class_name: 'ReturnRequest'
  has_many :Invoice, class_name: 'Invoice'

end
