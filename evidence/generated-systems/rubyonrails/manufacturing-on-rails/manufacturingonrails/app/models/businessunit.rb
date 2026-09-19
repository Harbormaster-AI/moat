class BusinessUnit < ApplicationRecord
  enum Category: [:ConsumerGoods, :IndustrialEquipment, :Electronics, :Pharmaceuticals, :FoodBeverage]


  has_many :Enterprise, class_name: 'Enterprise'
  has_many :Items, class_name: 'Item'
  has_many :Plants, class_name: 'Plant'

end
