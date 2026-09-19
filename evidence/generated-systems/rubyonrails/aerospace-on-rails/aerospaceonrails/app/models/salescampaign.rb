class SalesCampaign < ApplicationRecord
  enum Status: [:Prospecting, :Proposal, :Negotiation, :Won, :Lost]


  has_many :Region, class_name: 'SalesRegion'
  has_many :Operator, class_name: 'Operator'
  has_many :Quotes, class_name: 'Quote'

end
