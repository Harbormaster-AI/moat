class Exposure < ApplicationRecord
  enum ExposureType: [:BodilyInjury, :PropertyDamage, :Medical, :UninsuredMotorist, :PersonalInjuryProtection, :DwellingDamage, :ContentsDamage, :BusinessIncome]
  enum Status: [:Open, :Closed, :Pending, :Reserved]


  has_many :Claim, class_name: 'Claim'
  has_many :PolicyCoverage, class_name: 'PolicyCoverage'
  has_many :InsuredObject, class_name: 'InsuredObject'
  has_many :Reserves, class_name: 'ClaimReserve'
  has_many :Payments, class_name: 'ClaimPayment'

end
