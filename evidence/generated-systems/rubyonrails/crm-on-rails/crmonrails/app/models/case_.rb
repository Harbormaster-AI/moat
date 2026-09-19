class Case_ < ApplicationRecord
  enum Status: [:New, :Open, :PendingCustomer, :PendingExternal, :OnHold, :Resolved, :Closed, :Reopened]
  enum Priority: [:Low, :Medium, :High, :Critical]
  enum Origin: [:Email, :Phone, :Web, :Chat, :Social, :Community]
  enum Severity: [:Minor, :Major, :Critical, :Blocker]


  has_many :Organization, class_name: 'Organization'
  has_many :Account, class_name: 'Account'
  has_many :Contact, class_name: 'Contact'
  has_many :Owner, class_name: 'User'
  has_many :Team, class_name: 'Team'
  has_many :Activities, class_name: 'Activity'
  has_many :CaseComments, class_name: 'Note'
  has_many :Emails, class_name: 'EmailMessage'
  has_many :RelatedOpportunities, class_name: 'Opportunity'

end
