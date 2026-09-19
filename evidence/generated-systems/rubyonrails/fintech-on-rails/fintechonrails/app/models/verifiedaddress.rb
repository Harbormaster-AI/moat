class VerifiedAddress < ApplicationRecord
  enum VerificationStatus: [:Unverified, :Verified, :Failed]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :KycProfile, class_name: 'KYCProfile'

end
