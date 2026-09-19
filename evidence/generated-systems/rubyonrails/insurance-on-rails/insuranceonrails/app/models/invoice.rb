class Invoice < ApplicationRecord
  enum Status: [:Open, :Paid, :PartiallyPaid, :Void]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :BillingAccount, class_name: 'BillingAccount'
  has_many :Policy, class_name: 'Policy'
  has_many :Payments, class_name: 'Payment'

end
