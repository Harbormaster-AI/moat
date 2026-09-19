class Payout < ApplicationRecord
  enum Status: [:Pending, :Scheduled, :Paid, :Failed, :Cancelled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Seller, class_name: 'Seller'
  has_many :Orders, class_name: 'Order'

end
