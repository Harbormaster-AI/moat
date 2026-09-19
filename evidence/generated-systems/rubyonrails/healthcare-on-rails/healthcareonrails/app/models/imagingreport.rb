class ImagingReport < ApplicationRecord
  enum Status: [:Registered, :Partial, :Final, :Corrected, :Cancelled]


  has_many :ImagingOrder, class_name: 'ImagingOrder'
  has_many :Clinician, class_name: 'Clinician'
  has_many :Encounter, class_name: 'Encounter'
  has_many :ImagingCenter, class_name: 'ImagingCenter'

end
