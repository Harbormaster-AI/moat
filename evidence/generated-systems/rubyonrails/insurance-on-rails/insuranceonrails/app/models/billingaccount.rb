class BillingAccount < ApplicationRecord
  enum Status: [:Current, :Delinquent, :Collections, :Closed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Policies, class_name: 'Policy'
  has_many :Invoices, class_name: 'Invoice'
  has_many :Payments, class_name: 'Payment'

end
