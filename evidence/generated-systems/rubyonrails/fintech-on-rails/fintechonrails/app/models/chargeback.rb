class Chargeback < ApplicationRecord
  enum Stage: [:FirstChargeback, :SecondChargeback, :Arbitration]
  enum Status: [:Pending, :Accepted, :Reversed, :Lost]


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

  has_many :Dispute, class_name: 'Dispute'
  has_many :Transaction, class_name: 'Transaction'

end
