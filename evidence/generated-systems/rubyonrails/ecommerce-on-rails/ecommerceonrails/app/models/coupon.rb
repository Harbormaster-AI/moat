class Coupon < ApplicationRecord
  enum Status: [:Active, :Expired, :Disabled, :Exhausted]


  has_many :Promotion, class_name: 'Promotion'
  has_many :Redemptions, class_name: 'CouponRedemption'

end
