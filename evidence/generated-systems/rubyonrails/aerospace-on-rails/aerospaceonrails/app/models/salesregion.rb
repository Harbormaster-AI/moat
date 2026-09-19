class SalesRegion < ApplicationRecord


  has_many :Operators, class_name: 'Operator'
  has_many :SalesCampaigns, class_name: 'SalesCampaign'

end
