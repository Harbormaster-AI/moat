class Payment < ApplicationRecord
  enum Status: [:Authorized, :Captured, :PartiallyCaptured, :Declined, :Refunded, :PartiallyRefunded, :Voided, :Pending]
  enum PaymentMethod: [:CreditCard, :DebitCard, :PayPal, :BankTransfer, :CashOnDelivery, :GiftCard, :ApplePay, :GooglePay, :BuyNowPayLater]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Order, class_name: 'Order'
  has_many :Customer, class_name: 'Customer'
  has_many :PaymentProvider, class_name: 'PaymentProvider'
  has_many :Refunds, class_name: 'Refund'

end
