class PricingPlan < ApplicationRecord
  enum Status: [:Draft, :Active, :Suspended, :Archived]


  has_many :ProductOffering, class_name: 'ProductOffering'
  has_many :FeeSchedules, class_name: 'FeeSchedule'
  has_many :Limits, class_name: 'UsageLimit'

end
