class SalaryComponent < ApplicationRecord
  enum ComponentType: [:BaseSalary, :Allowance, :OvertimeRate, :Commission, :ShiftDifferential]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :CompensationPackage, class_name: 'CompensationPackage'

end
