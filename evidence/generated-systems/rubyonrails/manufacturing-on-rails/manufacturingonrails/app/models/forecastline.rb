class ForecastLine < ApplicationRecord


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Forecast, class_name: 'Forecast'
  has_many :Item, class_name: 'Item'

end
