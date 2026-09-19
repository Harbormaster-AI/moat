class EmailMessage < ApplicationRecord
  enum Direction: [:Inbound, :Outbound, :Internal]
  enum Status: [:Draft, :Sent, :Delivered, :Opened, :Bounced, :Failed, :Replied]


  has_many :Organization, class_name: 'Organization'
  has_many :Owner, class_name: 'User'
  has_many :Account, class_name: 'Account'
  has_many :Contact, class_name: 'Contact'
  has_many :Lead, class_name: 'Lead'
  has_many :Case, class_name: 'Case_'
  has_many :Opportunity, class_name: 'Opportunity'
  has_many :Campaign, class_name: 'Campaign'

end
