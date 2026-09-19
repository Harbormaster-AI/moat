class Subscription < ApplicationRecord
  enum Status: [:Active, :Paused, :Cancelled, :Expired]
  enum Interval: [:Weekly, :BiWeekly, :Monthly, :Quarterly, :SemiAnnual, :Annual]


  has_many :Customer, class_name: 'Customer'
  has_many :Variant, class_name: 'ProductVariant'
  has_many :PaymentProvider, class_name: 'PaymentProvider'
  has_many :Channel, class_name: 'Channel'

end
