class Opportunity < ApplicationRecord
  enum Stage: [:Qualification, :Discovery, :Proposal, :Negotiation, :ClosedWon, :ClosedLost]
  enum Type: [:NewBusiness, :ExistingBusiness, :Renewal, :Upsell, :CrossSell]
  enum ForecastCategory: [:Pipeline, :BestCase, :Commit, :Omitted, :Closed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :Account, class_name: 'Account'
  has_many :Owner, class_name: 'User'
  has_many :Contacts, class_name: 'Contact'
  has_many :LineItems, class_name: 'OpportunityLineItem'
  has_many :StageHistory, class_name: 'OpportunityStageHistory'
  has_many :Quotes, class_name: 'Quote'
  has_many :Orders, class_name: 'Order'
  has_many :Campaigns, class_name: 'Campaign'
  has_many :Activities, class_name: 'Activity'
  has_many :Teams, class_name: 'Team'

end
