class PaymentContract < ApplicationRecord
  enum Status: [:Draft, :Active, :Suspended, :Terminated]


  has_many :Merchant, class_name: 'Merchant'
  has_many :Acquirer, class_name: 'PaymentProcessor'

end
