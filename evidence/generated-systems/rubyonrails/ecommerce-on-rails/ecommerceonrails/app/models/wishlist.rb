class Wishlist < ApplicationRecord


  has_many :Customer, class_name: 'Customer'
  has_many :Items, class_name: 'WishlistItem'

end
