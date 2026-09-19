class Order < ApplicationRecord
  enum Status: [:Draft, :Submitted, :PartiallyFulfilled, :Fulfilled, :Invoiced, :Cancelled]


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

  has_many :Organization, class_name: 'Organization'
  has_many :Account, class_name: 'Account'
  has_many :Opportunity, class_name: 'Opportunity'
  has_many :Quote, class_name: 'Quote'
  has_many :Owner, class_name: 'User'
  has_many :Items, class_name: 'OrderItem'
  has_many :Contract, class_name: 'Contract'
  has_many :PriceBook, class_name: 'PriceBook'

end
