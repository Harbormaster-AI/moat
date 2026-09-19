class RepaymentSchedule < ApplicationRecord
  enum Status: [:Scheduled, :Paid, :Overdue, :Waived]


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

  has_many :Loan, class_name: 'Loan'
  has_many :Payments, class_name: 'Transaction'

end
