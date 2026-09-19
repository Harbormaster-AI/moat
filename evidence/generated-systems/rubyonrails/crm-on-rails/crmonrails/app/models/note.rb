class Note < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :Owner, class_name: 'User'
  has_many :Account, class_name: 'Account'
  has_many :Contact, class_name: 'Contact'
  has_many :Opportunity, class_name: 'Opportunity'
  has_many :Case, class_name: 'Case_'
  has_many :Lead, class_name: 'Lead'

end
