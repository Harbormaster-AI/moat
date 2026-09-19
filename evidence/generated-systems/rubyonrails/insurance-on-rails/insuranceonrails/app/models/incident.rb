class Incident < ApplicationRecord
  enum IncidentType: [:AutoAccident, :Fire, :Theft, :Windstorm, :Flood, :Hail, :Earthquake, :Vandalism, :Injury, :Death]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Claim, class_name: 'Claim'
  has_many :InsuredObjects, class_name: 'InsuredObject'

end
