class Position < ApplicationRecord


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

  has_many :Portfolio, class_name: 'InvestmentPortfolio'
  has_many :Security, class_name: 'Security'

end
