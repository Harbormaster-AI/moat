class Customer < ApplicationRecord
  enum CustomerGroup: [:Retail, :Wholesale, :VIP, :Employee]


  has_many :Addresses, class_name: 'CustomerAddress'
  has_many :Carts, class_name: 'Cart'
  has_many :Orders, class_name: 'Order'
  has_many :Payments, class_name: 'Payment'
  has_many :Reviews, class_name: 'Review'
  has_many :Wishlists, class_name: 'Wishlist'
  has_many :Subscriptions, class_name: 'Subscription'
  has_many :CouponRedemptions, class_name: 'CouponRedemption'
  has_many :GiftCards, class_name: 'GiftCard'

end
