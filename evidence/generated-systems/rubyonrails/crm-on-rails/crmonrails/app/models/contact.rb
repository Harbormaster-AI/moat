class Contact < ApplicationRecord
  enum PreferredContactMethod: [:Email, :Phone, :Mobile, :SMS, :InPerson, :Web]


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

  composed_of :phoneNumber,
    class_name: "PhoneNumber",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[phoneNumber_extension extension]
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

  has_many :Organization, class_name: 'Organization'
  has_many :Account, class_name: 'Account'
  has_many :Owner, class_name: 'User'
  has_many :Activities, class_name: 'Activity'
  has_many :Opportunities, class_name: 'Opportunity'
  has_many :Cases, class_name: 'Case_'
  has_many :Campaigns, class_name: 'Campaign'
  has_many :Notes, class_name: 'Note'
  has_many :EmailMessages, class_name: 'EmailMessage'

end
