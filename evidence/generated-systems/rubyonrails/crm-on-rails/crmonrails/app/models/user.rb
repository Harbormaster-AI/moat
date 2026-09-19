class User < ApplicationRecord
  enum Role: [:SalesRep, :SalesManager, :ServiceAgent, :MarketingSpecialist, :Administrator, :Executive]
  enum Status: [:Active, :Inactive, :Locked, :PendingInvite]


  composed_of :emailAddress,
    class_name: "EmailAddress",
    mapping: [
      %w[emailAddress_value value]
    ]

  composed_of :_Locale,
    class_name: "_Locale",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[_Locale_timezone timezone]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :Teams, class_name: 'Team'
  has_many :Activities, class_name: 'Activity'
  has_many :OwnedAccounts, class_name: 'Account'
  has_many :OwnedLeads, class_name: 'Lead'
  has_many :OwnedOpportunities, class_name: 'Opportunity'
  has_many :OwnedCases, class_name: 'Case_'
  has_many :Quotes, class_name: 'Quote'
  has_many :Orders, class_name: 'Order'
  has_many :Contracts, class_name: 'Contract'
  has_many :EmailMessages, class_name: 'EmailMessage'

end
