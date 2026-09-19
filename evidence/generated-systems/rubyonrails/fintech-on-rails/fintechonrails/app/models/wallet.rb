class Wallet < ApplicationRecord
  enum Status: [:Active, :Suspended, :Closed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Transactions, class_name: 'Transaction'

end
