class Payment < ApplicationRecord
  enum Method: [:ACH, :Check, :CreditCard, :EFT, :Cash]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Invoice, class_name: 'Invoice'
  has_many :Payer, class_name: 'InsurancePayer'

end
