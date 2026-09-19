class Category < ApplicationRecord


  has_many :Catalog, class_name: 'Catalog'
  has_many :ParentCategory, class_name: 'Category'
  has_many :Subcategories, class_name: 'Category'
  has_many :Products, class_name: 'Product'

end
