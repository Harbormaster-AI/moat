class Customer < ApplicationRecord
  enum CustomerType: [:Individual, :Organization]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Applications, class_name: 'Application'
  has_many :Policies, class_name: 'Policy'
  has_many :Claims, class_name: 'Claim'
  has_many :Agents, class_name: 'Agent'
  has_many :Beneficiaries, class_name: 'Beneficiary'

end
