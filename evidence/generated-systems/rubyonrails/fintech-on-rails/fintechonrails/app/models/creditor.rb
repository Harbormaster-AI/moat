class Creditor < ApplicationRecord


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

  has_many :Mandates, class_name: 'DirectDebitMandate'

end
