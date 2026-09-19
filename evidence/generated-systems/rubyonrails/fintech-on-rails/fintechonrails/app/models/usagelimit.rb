class UsageLimit < ApplicationRecord
  enum Scope: [:PerTransaction, :Daily, :Monthly, :Yearly, :Rolling24h]
  enum Period: [:None, :Day, :Week, :Month, :Year]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :PricingPlan, class_name: 'PricingPlan'

end
