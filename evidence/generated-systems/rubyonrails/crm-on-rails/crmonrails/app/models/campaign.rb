class Campaign < ApplicationRecord
  enum Status: [:Planned, :InProgress, :Completed, :OnHold, :Cancelled]
  enum Type: [:Email, :Social, :Event, :Webinar, :Advertising, :ContentMarketing, :Referral]


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

  has_many :Organization, class_name: 'Organization'
  has_many :ParentCampaign, class_name: 'Campaign'
  has_many :ChildCampaigns, class_name: 'Campaign'
  has_many :Members, class_name: 'CampaignMember'
  has_many :Opportunities, class_name: 'Opportunity'
  has_many :Accounts, class_name: 'Account'
  has_many :Leads, class_name: 'Lead'
  has_many :Contacts, class_name: 'Contact'
  has_many :Teams, class_name: 'Team'
  has_many :Activities, class_name: 'Activity'

end
