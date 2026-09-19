class LaboratoryOrder < ApplicationRecord
  enum SpecimenType: [:Blood, :Urine, :Saliva, :Sputum, :Tissue, :CSF, :Stool]


  has_many :Order, class_name: 'ClinicalOrder'
  has_many :Laboratory, class_name: 'Laboratory'
  has_many :Results, class_name: 'LabResult'

end
