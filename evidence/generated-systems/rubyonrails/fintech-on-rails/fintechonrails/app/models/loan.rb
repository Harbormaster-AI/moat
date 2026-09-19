class Loan < ApplicationRecord
  enum RateType: [:Fixed, :Variable]
  enum Status: [:Active, :Delinquent, :Closed, :ChargedOff]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Schedule, class_name: 'RepaymentSchedule'
  has_many :Collateral, class_name: 'Collateral'
  has_many :Transactions, class_name: 'LoanTransaction'

end
