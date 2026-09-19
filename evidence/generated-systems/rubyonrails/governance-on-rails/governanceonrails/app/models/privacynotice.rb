class PrivacyNotice < ApplicationRecord
  enum Status: [:Draft, :InReview, :Approved, :Retired]


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :Organization, class_name: 'Organization'
  has_many :Consents, class_name: 'Consent'

end
