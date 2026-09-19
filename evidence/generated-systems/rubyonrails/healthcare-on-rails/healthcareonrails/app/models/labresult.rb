class LabResult < ApplicationRecord
  enum Status: [:Registered, :Partial, :Final, :Corrected, :Cancelled]


  has_many :LaboratoryOrder, class_name: 'LaboratoryOrder'
  has_many :Observations, class_name: 'Observation'
  has_many :Laboratory, class_name: 'Laboratory'

end
