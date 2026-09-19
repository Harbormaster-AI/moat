class PayrollItem < ApplicationRecord
  enum ItemType: [:Earning, :Deduction, :Tax, :Benefit]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :PayrollRun, class_name: 'PayrollRun'
  has_many :Employee, class_name: 'Employee'

end
