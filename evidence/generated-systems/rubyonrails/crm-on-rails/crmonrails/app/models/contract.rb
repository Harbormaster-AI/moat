class Contract < ApplicationRecord
  enum Status: [:Draft, :Active, :Suspended, :Expired, :Terminated, :Renewed]


  has_many :Organization, class_name: 'Organization'
  has_many :Account, class_name: 'Account'
  has_many :Owner, class_name: 'User'
  has_many :Orders, class_name: 'Order'
  has_many :Cases, class_name: 'Case_'

end
