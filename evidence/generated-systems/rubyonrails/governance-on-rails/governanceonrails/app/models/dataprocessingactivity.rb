class DataProcessingActivity < ApplicationRecord
  enum LawfulBasis: [:Consent, :Contract, :LegalObligation, :VitalInterests, :PublicTask, :LegitimateInterests]


  has_many :Organization, class_name: 'Organization'
  has_many :DataCategories, class_name: 'DataCategory'
  has_many :Systems, class_name: 'System_'
  has_many :Records, class_name: 'Record_'
  has_many :PrivacyNotices, class_name: 'PrivacyNotice'
  has_many :ThirdParties, class_name: 'ThirdParty'
  has_many :Consents, class_name: 'Consent'
  has_many :DataBreaches, class_name: 'DataBreach'
  has_many :DataSubjectRequests, class_name: 'DataSubjectRequest'

end
