class RiskAssessment < ApplicationRecord
  enum AssessmentType: [:SelfAssessment, :InternalAssessment, :ExternalAssessment, :ReadinessReview]


  has_many :Risk, class_name: 'Risk'

end
