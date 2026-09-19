class Refund < ApplicationRecord
  enum Status: [:Requested, :Approved, :Declined, :Processed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Payment, class_name: 'Payment'
  has_many :Order, class_name: 'Order'

end
