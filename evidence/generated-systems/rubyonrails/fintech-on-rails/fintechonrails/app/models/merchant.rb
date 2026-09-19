class Merchant < ApplicationRecord


  has_many :Terminals, class_name: 'Terminal'
  has_many :PaymentContracts, class_name: 'PaymentContract'
  has_many :Payouts, class_name: 'Payout'
  has_many :Settlements, class_name: 'SettlementBatch'
  has_many :Disputes, class_name: 'Dispute'
  has_many :Invoices, class_name: 'Invoice'

end
