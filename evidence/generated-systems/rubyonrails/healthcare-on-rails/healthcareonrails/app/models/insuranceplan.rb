class InsurancePlan < ApplicationRecord
  enum PlanType: [:HMO, :PPO, :EPO, :POS, :Indemnity, :MedicareAdvantage, :MedicaidManagedCare]


  has_many :Payer, class_name: 'InsurancePayer'
  has_many :Coverages, class_name: 'Coverage'

end
