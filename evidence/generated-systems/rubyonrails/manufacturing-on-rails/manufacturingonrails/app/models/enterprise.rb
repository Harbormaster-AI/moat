class Enterprise < ApplicationRecord


  has_many :BusinessUnits, class_name: 'BusinessUnit'
  has_many :Plants, class_name: 'Plant'
  has_many :Suppliers, class_name: 'Supplier'
  has_many :Customers, class_name: 'Customer'

end
