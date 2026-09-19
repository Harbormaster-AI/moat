class FXDeal < ApplicationRecord
  enum Status: [:Booked, :Cancelled, :Settled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Quote, class_name: 'FXQuote'
  has_many :PaymentOrders, class_name: 'PaymentOrder'

end
