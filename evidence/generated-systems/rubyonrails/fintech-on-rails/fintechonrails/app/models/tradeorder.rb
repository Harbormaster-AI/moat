class TradeOrder < ApplicationRecord
  enum Side: [:Buy, :Sell]
  enum Type: [:Market, :Limit, :Stop, :StopLimit]
  enum Status: [:New, :PartiallyFilled, :Filled, :Cancelled, :Rejected, :Expired]
  enum TimeInForce: [:Day, :GTC, :IOC, :FOK]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Portfolio, class_name: 'InvestmentPortfolio'
  has_many :Security, class_name: 'Security'
  has_many :Trades, class_name: 'Trade'

end
