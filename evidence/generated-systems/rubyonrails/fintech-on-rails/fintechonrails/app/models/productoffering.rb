class ProductOffering < ApplicationRecord
  enum Category: [:Checking, :Savings, :CreditCard, :Loan, :Investment, :Insurance, :Payments, :FX, :Wallet]


  has_many :Institution, class_name: 'FinancialInstitution'
  has_many :PricingPlans, class_name: 'PricingPlan'

end
