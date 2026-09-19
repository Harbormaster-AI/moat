class SettlementBatch < ApplicationRecord
  enum Status: [:Open, :Processing, :Closed, :Reconciled]


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Processor, class_name: 'PaymentProcessor'
  has_many :Merchant, class_name: 'Merchant'
  has_many :Payouts, class_name: 'Payout'
  has_many :Transactions, class_name: 'Transaction'

end
