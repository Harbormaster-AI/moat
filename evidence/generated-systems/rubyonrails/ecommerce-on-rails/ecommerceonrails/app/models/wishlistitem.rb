class WishlistItem < ApplicationRecord


  has_many :Wishlist, class_name: 'Wishlist'
  has_many :Variant, class_name: 'ProductVariant'

end
