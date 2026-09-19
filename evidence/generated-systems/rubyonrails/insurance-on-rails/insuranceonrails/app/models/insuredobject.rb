class InsuredObject < ApplicationRecord
  enum ObjectType: [:Vehicle, :Property, :Person, :Equipment, :LiabilityExposure]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Policy, class_name: 'Policy'
  has_many :Coverages, class_name: 'PolicyCoverage'

end
