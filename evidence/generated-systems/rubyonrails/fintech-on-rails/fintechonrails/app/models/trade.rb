class Trade < ApplicationRecord


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
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

  has_many :Order, class_name: 'TradeOrder'
  has_many :Security, class_name: 'Security'
  has_many :InvestmentAccount, class_name: 'InvestmentAccount'

end
