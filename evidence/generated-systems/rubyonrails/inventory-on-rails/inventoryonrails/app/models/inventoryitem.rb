class InventoryItem < ApplicationRecord
  enum StockStatus: [:Available, :Reserved, :Damaged, :Hold, :Quarantined, :InTransit, :PendingInspection]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Location, class_name: 'StorageLocation'
  has_many :Lot, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'
  has_many :Transactions, class_name: 'InventoryTransaction'
  has_many :Reservations, class_name: 'Reservation'

end
