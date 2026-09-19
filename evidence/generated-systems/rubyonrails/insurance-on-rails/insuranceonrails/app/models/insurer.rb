class Insurer < ApplicationRecord


  has_many :Products, class_name: 'InsuranceProduct'
  has_many :DistributionPartners, class_name: 'Distributor'
  has_many :Policies, class_name: 'Policy'
  has_many :Claims, class_name: 'Claim'
  has_many :ReinsuranceAgreements, class_name: 'ReinsuranceAgreement'

end
