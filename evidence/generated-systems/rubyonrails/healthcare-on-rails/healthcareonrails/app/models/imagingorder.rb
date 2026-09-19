class ImagingOrder < ApplicationRecord
  enum Modality: [:XRay, :CT, :MRI, :Ultrasound, :PET, :Mammography]


  has_many :Order, class_name: 'ClinicalOrder'
  has_many :ImagingCenter, class_name: 'ImagingCenter'
  has_many :Reports, class_name: 'ImagingReport'

end
