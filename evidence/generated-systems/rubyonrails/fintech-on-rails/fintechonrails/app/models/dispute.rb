class Dispute < ApplicationRecord
  enum Reason: [:Fraud, :Duplicate, :NotAsDescribed, :NotReceived, :ProcessingError]
  enum Status: [:Open, :Represented, :Won, :Lost, :Closed]


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

  has_many :Transaction, class_name: 'Transaction'
  has_many :Card, class_name: 'PaymentCard'
  has_many :Merchant, class_name: 'Merchant'
  has_many :Chargebacks, class_name: 'Chargeback'

end
