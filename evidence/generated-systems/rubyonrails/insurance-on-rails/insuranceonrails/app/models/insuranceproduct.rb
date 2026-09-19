class InsuranceProduct < ApplicationRecord
  enum LineOfBusiness: [:PersonalAuto, :Homeowners, :Renters, :TermLife, :WholeLife, :CommercialProperty, :GeneralLiability, :WorkersCompensation]


  has_many :Insurer, class_name: 'Insurer'
  has_many :CoverageDefinitions, class_name: 'CoverageDefinition'

end
