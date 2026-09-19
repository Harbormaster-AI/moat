class CampaignMember < ApplicationRecord
  enum Status: [:Sent, :Opened, :Responded, :Unsubscribed, :Bounced, :Registered, :Attended, :NoShow]
  enum MemberType: [:Lead, :Contact]


  has_many :Campaign, class_name: 'Campaign'
  has_many :Lead, class_name: 'Lead'
  has_many :Contact, class_name: 'Contact'

end
