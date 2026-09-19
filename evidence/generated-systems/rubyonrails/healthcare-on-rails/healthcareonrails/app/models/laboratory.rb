class Laboratory < ApplicationRecord


  has_many :Facility, class_name: 'Facility'
  has_many :LaboratoryOrders, class_name: 'LaboratoryOrder'
  has_many :LabResults, class_name: 'LabResult'

end
