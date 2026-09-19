class InsurancePayer < ApplicationRecord
  enum PayerType: [:Commercial, :Government, :SelfInsured]


  has_many :Plans, class_name: 'InsurancePlan'
  has_many :Claims, class_name: 'Claim'

end
