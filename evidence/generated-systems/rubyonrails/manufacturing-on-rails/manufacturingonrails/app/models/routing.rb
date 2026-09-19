class Routing < ApplicationRecord
  enum RoutingType: [:Standard, :Alternate, :Rework]
  enum Status: [:Draft, :Released, :Obsolete]


  has_many :Item, class_name: 'Item'
  has_many :Operations, class_name: 'Operation'

end
