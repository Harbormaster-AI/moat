class BenefitPlan < ApplicationRecord
  enum BenefitType: [:Medical, :Dental, :Vision, :LifeInsurance, :Disability, :Retirement, :Wellness]


  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :Enrollments, class_name: 'BenefitEnrollment'

end
