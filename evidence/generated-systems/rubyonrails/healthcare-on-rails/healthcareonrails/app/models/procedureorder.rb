class ProcedureOrder < ApplicationRecord
  enum AnesthesiaType: [:None, :Local, :Regional, :General, :Sedation]


  has_many :Order, class_name: 'ClinicalOrder'
  has_many :Facility, class_name: 'Facility'
  has_many :Procedure, class_name: 'Procedure'

end
