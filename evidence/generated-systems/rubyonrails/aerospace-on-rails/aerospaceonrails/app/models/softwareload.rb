class SoftwareLoad < ApplicationRecord
  enum LoadType: [:FlightDeckSoftware, :MaintenanceTools, :CabinIFE, :ConnectivityModem]


  has_many :ConnectedAircraft, class_name: 'ConnectedAircraft'
  has_many :AvionicsSuite, class_name: 'AvionicsSuite'

end
