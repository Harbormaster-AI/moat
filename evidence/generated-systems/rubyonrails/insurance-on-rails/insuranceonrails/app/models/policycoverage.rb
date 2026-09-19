class PolicyCoverage < ApplicationRecord
  enum CoverageType: [:Liability, :Collision, :Comprehensive, :PropertyDamage, :BodilyInjury, :UninsuredMotorist, :MedicalPayments, :Dwelling, :Contents, :PersonalLiability, :BusinessInterruption, :ProfessionalLiability]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Policy, class_name: 'Policy'
  has_many :InsuredObjects, class_name: 'InsuredObject'

end
