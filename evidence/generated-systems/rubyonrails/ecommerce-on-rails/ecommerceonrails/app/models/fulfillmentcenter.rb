class FulfillmentCenter < ApplicationRecord


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Merchant, class_name: 'Merchant'
  has_many :InventoryItems, class_name: 'InventoryItem'
  has_many :Shipments, class_name: 'Shipment'

end
