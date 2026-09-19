class FeeSchedule < ApplicationRecord
  enum FeeType: [:Fixed, :Percentage, :Tiered, :Interchange, :Network, :Chargeback, :ATM, :FX]
  enum CalculationMethod: [:PerTransaction, :PerMonth, :PerAnnum, :Slab, :Tiered]


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

  has_many :PricingPlan, class_name: 'PricingPlan'

end
