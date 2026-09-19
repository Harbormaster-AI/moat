class Beneficiary < ApplicationRecord
  enum Relationship: [:Spouse, :Child, :Parent, :Sibling, :BusinessPartner, :Estate, :Trust, :Other]


  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Policy, class_name: 'Policy'
  has_many :Customer, class_name: 'Customer'

end
