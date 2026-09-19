class Payout < ApplicationRecord
  enum Status: [:Scheduled, :Processing, :Paid, :Failed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Merchant, class_name: 'Merchant'
  has_many :SettlementBatch, class_name: 'SettlementBatch'
  has_many :DestinationAccount, class_name: 'Account'

end
