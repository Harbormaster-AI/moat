class ThirdParty < ApplicationRecord
  enum PartyType: [:Individual, :Company, :GovernmentAgency]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Subrogations, class_name: 'SubrogationRecovery'

end
