class Consent < ApplicationRecord
  enum ConsentType: [:DataAccess, :PaymentInitiation]
  enum Status: [:Active, :Revoked, :Expired]


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

  has_many :Customer, class_name: 'Customer'
  has_many :ApiClient, class_name: 'APIClient'

end
