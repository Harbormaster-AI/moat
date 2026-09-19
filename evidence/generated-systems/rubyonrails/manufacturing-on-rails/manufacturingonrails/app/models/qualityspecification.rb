class QualitySpecification < ApplicationRecord


  has_many :Item, class_name: 'Item'

end
