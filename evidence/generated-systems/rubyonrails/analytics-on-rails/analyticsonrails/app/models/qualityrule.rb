class QualityRule < ApplicationRecord
  enum Dimension: [:Completeness, :Accuracy, :Consistency, :Timeliness, :Uniqueness, :Validity]
  enum Operator: [:GreaterThan, :GreaterThanOrEqual, :LessThan, :LessThanOrEqual, :Equal, :NotEqual]


  composed_of :threshold,
    class_name: "Threshold",
    mapping: [
      ${$mapping}, 
      %w[threshold_inclusive inclusive]
    ]

  has_many :Dataset, class_name: 'DataSet'
  has_many :Checks, class_name: 'QualityCheck'

end
