class Beneficiary < ApplicationRecord


  composed_of :accountIdentifier,
    class_name: "AccountIdentifier",
    mapping: [
      %w[accountIdentifier_value value]
    ]

  composed_of :iBAN,
    class_name: "IBAN",
    mapping: [
      %w[iBAN_value value]
    ]

  composed_of :bIC,
    class_name: "BIC",
    mapping: [
      %w[bIC_value value]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Customer, class_name: 'Customer'

end
