class LandingGear < ApplicationRecord
  enum GearType: [:Tricycle, :Tandem, :Taildragger, :Skid, :Floats, :Retractable]


  has_many :Supplier, class_name: 'Supplier'
  has_many :Variants, class_name: 'AircraftVariant'

end
