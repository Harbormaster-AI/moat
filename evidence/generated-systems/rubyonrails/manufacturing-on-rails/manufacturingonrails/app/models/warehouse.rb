class Warehouse < ApplicationRecord
  enum WarehouseType: [:RawMaterial, :WIP, :FinishedGoods, :Distribution]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Plant, class_name: 'Plant'
  has_many :Locations, class_name: 'Location'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
