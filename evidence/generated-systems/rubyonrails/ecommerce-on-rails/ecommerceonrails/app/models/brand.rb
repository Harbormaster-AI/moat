class Brand < ApplicationRecord


  has_many :Merchant, class_name: 'Merchant'
  has_many :Products, class_name: 'Product'

end
