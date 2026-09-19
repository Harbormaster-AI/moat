class TaxRule < ApplicationRecord
  enum TaxClass: [:Standard, :Reduced, :Zero, :Exempt, :DigitalServices, :Food, :Clothing]


  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Merchant, class_name: 'Merchant'
  has_many :Channels, class_name: 'Channel'

end
