class QuoteLineItem < ApplicationRecord


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

  has_many :Quote, class_name: 'Quote'
  has_many :Product, class_name: 'Product'
  has_many :PriceBookEntry, class_name: 'PriceBookEntry'
  has_many :OpportunityLineItem, class_name: 'OpportunityLineItem'

end
