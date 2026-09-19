class TransferOrder < ApplicationRecord
  enum Status: [:Draft, :Released, :InTransit, :Received, :Closed, :Cancelled]


  has_many :OriginWarehouse, class_name: 'Warehouse'
  has_many :DestinationWarehouse, class_name: 'Warehouse'
  has_many :Lines, class_name: 'TransferOrderLine'
  has_many :Transactions, class_name: 'InventoryTransaction'

end
