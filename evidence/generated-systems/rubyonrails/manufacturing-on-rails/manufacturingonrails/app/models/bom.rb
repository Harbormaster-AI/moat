class BOM < ApplicationRecord
  enum Status: [:Draft, :Released, :Obsolete]


  has_many :ParentItem, class_name: 'Item'
  has_many :BomItems, class_name: 'BOMItem'

end
