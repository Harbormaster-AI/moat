class InboundShipment < ApplicationRecord
  enum Status: [:Planned, :Arrived, :Received, :Closed, :Cancelled]


  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Lines, class_name: 'InboundShipmentLine'
  has_many :Transactions, class_name: 'InventoryTransaction'

end
