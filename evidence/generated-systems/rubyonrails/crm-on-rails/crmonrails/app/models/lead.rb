class Lead < ApplicationRecord
  enum Status: [:New, :Working, :Nurturing, :Qualified, :Unqualified, :Converted]
  enum Source: [:Web, :Referral, :Event, :Partner, :Advertisement, :Outbound, :Inbound, :Social, :Other]
  enum Rating: [:Hot, :Warm, :Cold]


  composed_of :emailAddress,
    class_name: "EmailAddress",
    mapping: [
      %w[emailAddress_value value]
    ]

  composed_of :phoneNumber,
    class_name: "PhoneNumber",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[phoneNumber_extension extension]
    ]

  has_many :Organization, class_name: 'Organization'
  has_many :Owner, class_name: 'User'
  has_many :Activities, class_name: 'Activity'
  has_many :Campaigns, class_name: 'Campaign'
  has_many :ConvertedAccount, class_name: 'Account'
  has_many :ConvertedContact, class_name: 'Contact'
  has_many :ConvertedOpportunity, class_name: 'Opportunity'
  has_many :Notes, class_name: 'Note'
  has_many :EmailMessages, class_name: 'EmailMessage'

end
