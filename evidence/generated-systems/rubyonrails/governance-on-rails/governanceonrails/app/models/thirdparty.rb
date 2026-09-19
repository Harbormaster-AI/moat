class ThirdParty < ApplicationRecord
  enum ThirdPartyType: [:Vendor, :Processor, :JointController, :Subprocessor, :Partner, :Consultant]
  enum Criticality: [:Low, :Medium, :High, :Critical]


  composed_of :emailAddress,
    class_name: "EmailAddress",
    mapping: [
      %w[emailAddress_value value]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :Assessments, class_name: 'ThirdPartyAssessment'
  has_many :Contracts, class_name: 'Contract'
  has_many :Obligations, class_name: 'Obligation'
  has_many :DataBreaches, class_name: 'DataBreach'

end
