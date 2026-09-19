class FXQuote < ApplicationRecord
  enum PriceType: [:Indicative, :Firm]


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

  has_many :RequestedBy, class_name: 'Customer'

end
