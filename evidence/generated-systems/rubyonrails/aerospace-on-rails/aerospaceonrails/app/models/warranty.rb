class Warranty < ApplicationRecord
  enum WarrantyType: [:Basic, :Powerplant, :Avionics, :Corrosion]


  has_many :Aircraft, class_name: 'Aircraft'

end
