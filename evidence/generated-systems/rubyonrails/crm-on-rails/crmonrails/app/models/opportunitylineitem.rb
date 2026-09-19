class OpportunityLineItem < ApplicationRecord


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Opportunity, class_name: 'Opportunity'
  has_many :Product, class_name: 'Product'
  has_many :PriceBookEntry, class_name: 'PriceBookEntry'

end
