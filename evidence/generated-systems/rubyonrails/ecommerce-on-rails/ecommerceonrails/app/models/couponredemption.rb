class CouponRedemption < ApplicationRecord


  has_many :Coupon, class_name: 'Coupon'
  has_many :Order, class_name: 'Order'
  has_many :Customer, class_name: 'Customer'

end
