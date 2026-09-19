class Catalog < ApplicationRecord


  has_many :Channel, class_name: 'Channel'
  has_many :Categories, class_name: 'Category'

end
