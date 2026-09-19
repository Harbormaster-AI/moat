class Collateral < ApplicationRecord
  enum CollateralType: [:RealEstate, :Deposit, :PersonalGuarantee, :Inventory, :Equipment]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Loan, class_name: 'Loan'

end
