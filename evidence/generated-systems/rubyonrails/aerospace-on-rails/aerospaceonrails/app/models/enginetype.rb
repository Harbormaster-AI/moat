class EngineType < ApplicationRecord
  enum Category: [:Turbofan, :Turboprop, :Turbojet, :Piston, :Electric, :Rocket]


  has_many :Supplier, class_name: 'Supplier'
  has_many :CompatibleModels, class_name: 'AircraftModel'

end
