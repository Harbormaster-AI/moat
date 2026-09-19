class FinancialInstitution < ApplicationRecord


  composed_of :bIC,
    class_name: "BIC",
    mapping: [
      %w[bIC_value value]
    ]

  has_many :Branches, class_name: 'Branch'
  has_many :Customers, class_name: 'Customer'
  has_many :ProductOfferings, class_name: 'ProductOffering'
  has_many :PaymentProcessors, class_name: 'PaymentProcessor'
  has_many :CompliancePolicies, class_name: 'CompliancePolicy'

end
