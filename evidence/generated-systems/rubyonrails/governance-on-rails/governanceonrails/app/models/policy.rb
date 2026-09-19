class Policy < ApplicationRecord
  enum PolicyType: [:InformationSecurity, :DataProtection, :Ethics, :RecordsManagement, :RiskManagement, :Compliance, :Privacy, :AcceptableUse]
  enum Status: [:Draft, :InReview, :Approved, :Retired]


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :Owners, class_name: 'Person'
  has_many :RelatedRequirements, class_name: 'ComplianceRequirement'
  has_many :Controls, class_name: 'Control'
  has_many :Procedures, class_name: 'Procedure'
  has_many :Exceptions, class_name: 'Exception_'
  has_many :Attestations, class_name: 'Attestation'

end
