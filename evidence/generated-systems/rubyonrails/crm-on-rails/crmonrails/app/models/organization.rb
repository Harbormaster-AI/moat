class Organization < ApplicationRecord


  composed_of :_Locale,
    class_name: "_Locale",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[_Locale_timezone timezone]
    ]

  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :Users, class_name: 'User'
  has_many :Accounts, class_name: 'Account'
  has_many :Teams, class_name: 'Team'
  has_many :Territories, class_name: 'Territory'
  has_many :Products, class_name: 'Product'
  has_many :PriceBooks, class_name: 'PriceBook'
  has_many :Campaigns, class_name: 'Campaign'

end
