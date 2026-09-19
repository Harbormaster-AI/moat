class TaxWithholding < ApplicationRecord
  enum FilingStatus: [:Single, :MarriedFilingJointly, :MarriedFilingSeparately, :HeadOfHousehold, :QualifyingWidowEr]


  composed_of :taxId,
    class_name: "TaxId",
    mapping: [
      ${$mapping}, 
      %w[taxId_country country]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Employee, class_name: 'Employee'

end
