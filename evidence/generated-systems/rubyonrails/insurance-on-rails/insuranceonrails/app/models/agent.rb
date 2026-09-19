class Agent < ApplicationRecord
  enum Status: [:Active, :Suspended, :Terminated]


  has_many :Distributor, class_name: 'Distributor'
  has_many :Policies, class_name: 'Policy'
  has_many :Customers, class_name: 'Customer'

end
