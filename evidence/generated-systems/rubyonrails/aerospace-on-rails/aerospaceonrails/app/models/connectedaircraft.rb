class ConnectedAircraft < ApplicationRecord
  enum ConnectivityStatus: [:Offline, :Online, :Degraded]


  has_many :Aircraft, class_name: 'Aircraft'
  has_many :FlightHealthEvents, class_name: 'FlightHealthEvent'
  has_many :SoftwareLoads, class_name: 'SoftwareLoad'

end
