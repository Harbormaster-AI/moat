class InvestmentAccount < ApplicationRecord
  enum AccountType: [:Brokerage, :Retirement, :Custody, :Margin]
  enum Status: [:Pending, :Active, :Frozen, :Closed]


  composed_of :accountNumber,
    class_name: "AccountNumber",
    mapping: [
      %w[accountNumber_value value]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Portfolio, class_name: 'InvestmentPortfolio'
  has_many :Trades, class_name: 'Trade'
  has_many :Orders, class_name: 'TradeOrder'

end
