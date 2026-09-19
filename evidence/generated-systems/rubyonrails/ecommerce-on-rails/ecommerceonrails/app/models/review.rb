class Review < ApplicationRecord
  enum Status: [:Pending, :Approved, :Rejected, :Flagged]


  has_many :Product, class_name: 'Product'
  has_many :Customer, class_name: 'Customer'
  has_many :Order, class_name: 'Order'

end
