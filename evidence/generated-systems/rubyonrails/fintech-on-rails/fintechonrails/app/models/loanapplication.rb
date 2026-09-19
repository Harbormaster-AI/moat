class LoanApplication < ApplicationRecord
  enum Product: [:PersonalLoan, :Mortgage, :InstallmentLoan, :CreditLine, :SME]
  enum Purpose: [:HomeImprovement, :Education, :DebtConsolidation, :Business, :Other]
  enum Status: [:Draft, :Submitted, :Underwriting, :Approved, :Declined, :Withdrawn]


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

  has_many :Customer, class_name: 'Customer'
  has_many :RiskAssessment, class_name: 'RiskAssessment'
  has_many :Loan, class_name: 'Loan'

end
