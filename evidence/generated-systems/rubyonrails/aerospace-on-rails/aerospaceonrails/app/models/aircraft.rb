class Aircraft < ApplicationRecord


  composed_of :mSN,
    class_name: "MSN",
    mapping: [
      %w[mSN_value value]
    ]

  has_many :Variant, class_name: 'AircraftVariant'
  has_many :Operator, class_name: 'Operator'
  has_many :Registration, class_name: 'Registration'
  has_many :Warranty, class_name: 'Warranty'
  has_many :MaintenanceRecords, class_name: 'MaintenanceWorkOrder'
  has_many :ConnectedAircraft, class_name: 'ConnectedAircraft'
  has_many :CabinLayout, class_name: 'CabinLayout'

end
