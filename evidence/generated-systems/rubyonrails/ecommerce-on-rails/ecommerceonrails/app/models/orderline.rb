class OrderLine < ApplicationRecord
  enum LineStatus: [:Pending, :Fulfilled, :Cancelled, :Backordered, :Returned]


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

  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Order, class_name: 'Order'
  has_many :Variant, class_name: 'ProductVariant'
  has_many :AppliedPromotions, class_name: 'Promotion'

end
