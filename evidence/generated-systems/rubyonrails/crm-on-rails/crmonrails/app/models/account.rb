class Account < ApplicationRecord
  enum AccountType: [:Prospect, :Customer, :Partner, :Vendor, :Competitor]
  enum LifecycleStage: [:Subscriber, :Lead, :MarketingQualified, :SalesQualified, :Customer, :Evangelist, :Churned]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  composed_of :phoneNumber,
    class_name: "PhoneNumber",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[phoneNumber_extension extension]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :ParentAccount, class_name: 'Account'
  has_many :ChildAccounts, class_name: 'Account'
  has_many :Contacts, class_name: 'Contact'
  has_many :Opportunities, class_name: 'Opportunity'
  has_many :Cases, class_name: 'Case_'
  has_many :Owner, class_name: 'User'
  has_many :Territory, class_name: 'Territory'
  has_many :Activities, class_name: 'Activity'
  has_many :Campaigns, class_name: 'Campaign'
  has_many :Quotes, class_name: 'Quote'
  has_many :Orders, class_name: 'Order'
  has_many :Contracts, class_name: 'Contract'
  has_many :Notes, class_name: 'Note'
  has_many :EmailMessages, class_name: 'EmailMessage'

end
