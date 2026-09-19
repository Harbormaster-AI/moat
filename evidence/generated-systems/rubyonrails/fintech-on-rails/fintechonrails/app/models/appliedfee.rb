class AppliedFee < ApplicationRecord
  enum FeeType: [:Fixed, :Percentage, :Tiered, :Interchange, :Network, :Chargeback, :ATM, :FX]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :PaymentOrder, class_name: 'PaymentOrder'
  has_many :Transaction, class_name: 'Transaction'

end
