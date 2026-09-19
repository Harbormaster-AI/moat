class Warehouse < ApplicationRecord


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :StorageLocations, class_name: 'StorageLocation'
  has_many :InventoryItems, class_name: 'InventoryItem'
  has_many :InboundShipments, class_name: 'InboundShipment'
  has_many :OutboundAllocations, class_name: 'OutboundAllocation'
  has_many :OriginTransfers, class_name: 'TransferOrder'
  has_many :DestinationTransfers, class_name: 'TransferOrder'
  has_many :CycleCounts, class_name: 'CycleCount'

end
