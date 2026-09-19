class SubrogationRecovery < ApplicationRecord
  enum Status: [:Open, :Negotiating, :Settled, :Uncollectible, :Closed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Claim, class_name: 'Claim'
  has_many :Exposure, class_name: 'Exposure'
  has_many :Counterparty, class_name: 'ThirdParty'

end
