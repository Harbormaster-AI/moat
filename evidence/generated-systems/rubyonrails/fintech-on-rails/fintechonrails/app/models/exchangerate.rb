class ExchangeRate < ApplicationRecord


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :UsedByQuotes, class_name: 'FXQuote'

end
