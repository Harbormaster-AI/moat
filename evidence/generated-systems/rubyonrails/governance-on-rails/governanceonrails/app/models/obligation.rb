class Obligation < ApplicationRecord
  enum ObligationType: [:Regulatory, :Contractual, :PolicyDerived, :IndustryStandard]
  enum ReviewFrequency: [:Continuous, :Daily, :Weekly, :Monthly, :Quarterly, :Annually, :AdHoc]


  has_many :Regulation, class_name: 'Regulation'
  has_many :Controls, class_name: 'Control'
  has_many :Policies, class_name: 'Policy'
  has_many :Contracts, class_name: 'Contract'

end
