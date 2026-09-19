class Quote < ApplicationRecord
  enum Status: [:Draft, :Presented, :Approved, :Rejected, :Accepted, :Expired, :Withdrawn]


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
  has_many :Owner, class_name: 'User'
  has_many :LineItems, class_name: 'QuoteLineItem'
  has_many :PriceBook, class_name: 'PriceBook'
  has_many :Order, class_name: 'Order'

end
