class Endorsement < ApplicationRecord


  has_many :Policy, class_name: 'Policy'

end
