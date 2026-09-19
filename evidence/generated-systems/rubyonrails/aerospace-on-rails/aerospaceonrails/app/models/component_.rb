class Component_ < ApplicationRecord
  enum ComponentCategory: [:Structure, :System, :Avionics, :Interior, :LandingGear, :Powerplant, :Consumable]
  enum SerializationMethod: [:Serialized, :LotTracked, :None]


  has_many :Supplier, class_name: 'Supplier'

end
