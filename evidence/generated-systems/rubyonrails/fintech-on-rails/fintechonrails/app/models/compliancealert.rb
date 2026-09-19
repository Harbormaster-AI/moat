class ComplianceAlert < ApplicationRecord
  enum Severity: [:Low, :Medium, :High, :Critical]
  enum Status: [:Open, :Investigating, :Resolved, :Dismissed]


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Screening, class_name: 'Screening'
  has_many :Transaction, class_name: 'Transaction'

end
