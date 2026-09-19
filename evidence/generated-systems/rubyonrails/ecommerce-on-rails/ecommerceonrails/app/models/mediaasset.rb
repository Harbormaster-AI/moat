class MediaAsset < ApplicationRecord
  enum MediaType: [:Image, :Video, :Document, :Audio, :Other]


  has_many :Product, class_name: 'Product'
  has_many :Variant, class_name: 'ProductVariant'

end
