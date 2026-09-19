class CompensationPackage < ApplicationRecord


  has_many :Contract, class_name: 'EmploymentContract'
  has_many :SalaryComponents, class_name: 'SalaryComponent'
  has_many :BonusPlans, class_name: 'BonusPlan'
  has_many :EquityGrants, class_name: 'EquityGrant'

end
