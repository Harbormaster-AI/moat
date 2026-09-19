class Transaction < ApplicationRecord
  enum TransactionType: [:Deposit, :Withdrawal, :Transfer, :Payment, :Refund, :Fee, :Interest, :FXConversion]
  enum Status: [:Pending, :Authorized, :Posted, :Settled, :Reversed, :Failed]


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

  has_many :Account, class_name: 'Account'
  has_many :Wallet, class_name: 'Wallet'
  has_many :PaymentOrder, class_name: 'PaymentOrder'
  has_many :Merchant, class_name: 'Merchant'
  has_many :Card, class_name: 'PaymentCard'
  has_many :RelatedTransactions, class_name: 'Transaction'
  has_many :Alerts, class_name: 'ComplianceAlert'

end
