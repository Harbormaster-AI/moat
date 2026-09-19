class Adjuster < ApplicationRecord
  enum AdjusterType: [:Staff, :Independent, :Public]


  has_many :Claims, class_name: 'Claim'
  has_many :ServiceProviders, class_name: 'ServiceProvider'

end
