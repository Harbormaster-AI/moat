class Product < ApplicationRecord
  enum ProductType: [:Good, :Service, :Subscription, :Bundle]
  enum Uom: [:Each, :Hour, :Day, :Month, :User, :Package]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :PriceBookEntries, class_name: 'PriceBookEntry'
  has_many :OpportunityLineItems, class_name: 'OpportunityLineItem'
  has_many :QuoteLineItems, class_name: 'QuoteLineItem'
  has_many :OrderItems, class_name: 'OrderItem'

end
