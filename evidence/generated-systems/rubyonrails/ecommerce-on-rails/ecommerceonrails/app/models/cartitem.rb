class CartItem < ApplicationRecord


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

  has_many :Cart, class_name: 'Cart'
  has_many :Variant, class_name: 'ProductVariant'
  has_many :AppliedPromotions, class_name: 'Promotion'

end
