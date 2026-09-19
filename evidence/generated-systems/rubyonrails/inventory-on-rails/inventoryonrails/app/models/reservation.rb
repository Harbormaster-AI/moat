class Reservation < ApplicationRecord
  enum ReservationStatus: [:Draft, :Confirmed, :Released, :Fulfilled, :Cancelled, :Expired]
  enum ReservationType: [:SalesOrder, :WorkOrder, :TransferOrder, :ServiceOrder, :Other]


  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Location, class_name: 'StorageLocation'
  has_many :InventoryItem, class_name: 'InventoryItem'
  has_many :Lot, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'
  has_many :DemandSignal, class_name: 'DemandSignal'

end
