class Plant < ApplicationRecord


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Manufacturer, class_name: 'AerospaceManufacturer'
  has_many :ProductionLines, class_name: 'ProductionLine'
  has_many :Warehouses, class_name: 'Warehouse'

end
