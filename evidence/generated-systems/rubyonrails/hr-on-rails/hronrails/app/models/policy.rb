class Policy < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :Acknowledgements, class_name: 'PolicyAcknowledgement'

end
