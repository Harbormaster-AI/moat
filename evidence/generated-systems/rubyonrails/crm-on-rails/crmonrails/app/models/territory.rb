class Territory < ApplicationRecord
  enum TerritoryType: [:Geographic, :Industry, :NamedAccount, :Segment, :Hybrid]


  has_many :Organization, class_name: 'Organization'
  has_many :Accounts, class_name: 'Account'
  has_many :Users, class_name: 'User'

end
