class InvestmentPortfolio < ApplicationRecord
  enum Status: [:Active, :Closed, :Suspended]


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Accounts, class_name: 'InvestmentAccount'
  has_many :Orders, class_name: 'TradeOrder'
  has_many :Holdings, class_name: 'Position'

end
