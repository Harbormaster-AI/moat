class Consent < ApplicationRecord
  enum ConsentType: [:Marketing, :Profiling, :Cookies, :Location, :Biometric]
  enum Status: [:Granted, :Withdrawn, :Expired, :NotRequired]


  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :PrivacyNotice, class_name: 'PrivacyNotice'

end
