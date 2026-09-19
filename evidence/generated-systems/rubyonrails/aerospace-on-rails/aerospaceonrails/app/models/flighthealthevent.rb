class FlightHealthEvent < ApplicationRecord
  enum Severity: [:Info, :Warning, :Critical]


  has_many :ConnectedAircraft, class_name: 'ConnectedAircraft'

end
