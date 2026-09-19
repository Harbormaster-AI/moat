class GiftCard < ApplicationRecord
  enum Status: [:Active, :Redeemed, :Expired, :Disabled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :IssuedOrder, class_name: 'Order'
  has_many :Redemptions, class_name: 'GiftCardRedemption'

end
