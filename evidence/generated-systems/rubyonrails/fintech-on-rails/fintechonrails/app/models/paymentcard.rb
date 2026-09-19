class PaymentCard < ApplicationRecord
  enum Scheme: [:Visa, :Mastercard, :Amex, :Discover, :UnionPay]
  enum Status: [:Active, :Blocked, :Closed, :Expired]


  composed_of :cardNumberToken,
    class_name: "CardNumberToken",
    mapping: [
      %w[cardNumberToken_value value]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Account, class_name: 'Account'
  has_many :Tokenizations, class_name: 'CardTokenization'
  has_many :Disputes, class_name: 'Dispute'

end
