class ImagingCenter < ApplicationRecord


  has_many :Facility, class_name: 'Facility'
  has_many :ImagingOrders, class_name: 'ImagingOrder'
  has_many :ImagingReports, class_name: 'ImagingReport'

end
