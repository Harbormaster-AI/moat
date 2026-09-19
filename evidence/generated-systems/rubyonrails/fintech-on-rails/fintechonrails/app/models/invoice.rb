class Invoice < ApplicationRecord
  enum Status: [:Draft, :Issued, :Paid, :Overdue, :Cancelled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Merchant, class_name: 'Merchant'
  has_many :Payments, class_name: 'PaymentOrder'

end
