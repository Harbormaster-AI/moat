class Payment < ApplicationRecord
  enum Method: [:ACH, :CreditCard, :DebitCard, :Check, :Cash, :Wire]
  enum Status: [:Pending, :Settled, :Failed, :Refunded, :Reversed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Invoice, class_name: 'Invoice'
  has_many :BillingAccount, class_name: 'BillingAccount'
  has_many :Policy, class_name: 'Policy'

end
