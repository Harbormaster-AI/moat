class ThirdPartyAssessment < ApplicationRecord
  enum AssessmentType: [:SelfAssessment, :InternalAssessment, :ExternalAssessment, :ReadinessReview]
  enum Result: [:Pass, :ConditionalPass, :Fail]


  has_many :ThirdParty, class_name: 'ThirdParty'
  has_many :Issues, class_name: 'Issue'

end
