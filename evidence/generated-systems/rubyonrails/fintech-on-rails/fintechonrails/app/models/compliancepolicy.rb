class CompliancePolicy < ApplicationRecord
  enum Status: [:Draft, :Active, :Retired]


  has_many :Institution, class_name: 'FinancialInstitution'

end
