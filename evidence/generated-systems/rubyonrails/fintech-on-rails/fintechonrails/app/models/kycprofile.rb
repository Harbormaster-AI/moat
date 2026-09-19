class KYCProfile < ApplicationRecord
  enum Status: [:Pending, :Verified, :Rejected, :Expired]
  enum VerificationLevel: [:Basic, :Standard, :Enhanced]


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Documents, class_name: 'KYCDocument'
  has_many :Screenings, class_name: 'Screening'
  has_many :Addresses, class_name: 'VerifiedAddress'

end
