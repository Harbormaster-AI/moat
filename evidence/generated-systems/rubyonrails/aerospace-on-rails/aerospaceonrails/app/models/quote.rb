class Quote < ApplicationRecord


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :AircraftOrder, class_name: 'AircraftOrder'

end
