class PriceBookEntry < ApplicationRecord


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :PriceBook, class_name: 'PriceBook'
  has_many :Product, class_name: 'Product'

end
