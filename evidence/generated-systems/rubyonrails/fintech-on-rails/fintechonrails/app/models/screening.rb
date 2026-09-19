class Screening < ApplicationRecord
  enum ScreeningType: [:Sanctions, :PEP, :AdverseMedia]
  enum Status: [:Clear, :Review, :Match]


  composed_of :riskScore,
    class_name: "RiskScore",
    mapping: [
      %w[riskScore_value value]
    ]

  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :KycProfile, class_name: 'KYCProfile'
  has_many :Alerts, class_name: 'ComplianceAlert'

end
