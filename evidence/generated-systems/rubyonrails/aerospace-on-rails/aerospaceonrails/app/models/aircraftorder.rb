class AircraftOrder < ApplicationRecord
  enum Status: [:Draft, :Committed, :InProduction, :Delivered, :Cancelled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Operator, class_name: 'Operator'
  has_many :Variant, class_name: 'AircraftVariant'
  has_many :Quote, class_name: 'Quote'
  has_many :PurchaseAgreement, class_name: 'PurchaseAgreement'

end
