class RiskAssessment < ApplicationRecord
  enum Decision: [:Approve, :Decline, :Refer]


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

  has_many :Application, class_name: 'LoanApplication'

end
