class ClaimPayment < ApplicationRecord
  enum PayeeType: [:Claimant, :Beneficiary, :ServiceProvider, :Lienholder, :Attorney]
  enum Method: [:ACH, :CreditCard, :DebitCard, :Check, :Cash, :Wire]
  enum Status: [:Pending, :Settled, :Failed, :Refunded, :Reversed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Claim, class_name: 'Claim'
  has_many :Exposure, class_name: 'Exposure'
  has_many :Beneficiary, class_name: 'Beneficiary'
  has_many :ServiceProvider, class_name: 'ServiceProvider'
  has_many :Customer, class_name: 'Customer'

end
