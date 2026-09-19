class Security < ApplicationRecord
  enum SecurityType: [:Equity, :Bond, :ETF, :MutualFund, :Derivative, :Crypto]


  has_many :Positions, class_name: 'Position'
  has_many :Trades, class_name: 'Trade'
  has_many :Orders, class_name: 'TradeOrder'

end
