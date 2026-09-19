class KYCDocument < ApplicationRecord
  enum DocumentType: [:Passport, :NationalID, :DriverLicense, :BusinessRegistration, :ProofOfAddress]
  enum Status: [:Submitted, :Approved, :Rejected, :Expired]


  composed_of :documentReference,
    class_name: "DocumentReference",
    mapping: [
      %w[documentReference_value value]
    ]

  has_many :KycProfile, class_name: 'KYCProfile'

end
