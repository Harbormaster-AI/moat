class GiftCardRedemption < ApplicationRecord


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :GiftCard, class_name: 'GiftCard'
  has_many :Order, class_name: 'Order'

end
