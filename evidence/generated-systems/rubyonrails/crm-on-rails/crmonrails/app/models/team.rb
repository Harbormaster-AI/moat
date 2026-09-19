class Team < ApplicationRecord
  enum TeamType: [:Sales, :Service, :Marketing, :AccountTeam, :DealDesk, :CrossFunctional]


  has_many :Organization, class_name: 'Organization'
  has_many :Users, class_name: 'User'
  has_many :Accounts, class_name: 'Account'
  has_many :Opportunities, class_name: 'Opportunity'
  has_many :Cases, class_name: 'Case_'
  has_many :Campaigns, class_name: 'Campaign'

end
