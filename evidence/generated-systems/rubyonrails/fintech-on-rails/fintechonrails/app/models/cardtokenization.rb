class CardTokenization < ApplicationRecord
  enum WalletProvider: [:ApplePay, :GooglePay, :SamsungPay, :Other]
  enum Status: [:Active, :Suspended, :Deactivated]


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Card, class_name: 'PaymentCard'

end
