class PriceBook < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :Entries, class_name: 'PriceBookEntry'
  has_many :Quotes, class_name: 'Quote'
  has_many :Orders, class_name: 'Order'

end
